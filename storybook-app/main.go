package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"

	. "storybook-app/helper"
	"storybook-app/router"
)

var serviceName string
var serviceTimestamp string

func main() {
	serviceName = "storybook-app"
	serviceAddr := "localhost:6010"
	serviceTimestamp = fmt.Sprintf("%v", time.Now().UnixMilli())

	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lshortfile)

	httpService := &http.Server{
		Addr:    serviceAddr,
		Handler: http.HandlerFunc(handler),
	}

	// https://dev.to/mokiat/proper-http-shutdown-in-go-3fji
	log.Println(fmt.Sprintf("http service started, serviceAddr=%s", serviceAddr))
	startupHttp(httpService)
	shutdownHttpHook(httpService)
	log.Println(fmt.Sprintf("http service stopped"))
}

func startupHttp(service *http.Server) {
	go func() {
		err := service.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http startup error: %v", err)
		}
	}()
}

func shutdownHttpHook(service *http.Server) {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	<-channel

	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := service.Shutdown(ctx)
	if err != nil {
		log.Fatalf("http shutdown error: %v", err)
	}
}

func handler(writer http.ResponseWriter, req *http.Request) {
	switch {
	case req.Method == "GET" && req.URL.Path == "/site":
		handlerSite(writer, req)
	case req.Method == "GET" && req.URL.Path == "/site-refresh":
		handlerSiteRefresh(writer, req)
	default:
		handlerIndex(writer, req)
	}
}

func handlerIndex(writer http.ResponseWriter, req *http.Request) {
	pathParam := req.URL.Path
	pageParam := ""
	if strings.HasPrefix(pathParam, "/pages/") {
		pageParam = strings.TrimPrefix(pathParam, "/pages/")
	}

	writer.Header().Set("content-type", "text/html; charset=utf-8")
	_ = Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(g.Attr("charset", "utf-8")),
				TitleEl(g.Text(serviceName)),
				Meta(
					// CSP policy for htmx server content to mitigate cross-site scripting.
					g.Attr("http-equiv", "content-security-policy"),
					g.Attr("content", "default-src 'self' 'unsafe-inline' 'unsafe-eval' unpkg.com;"),
				),
				Script(Src("https://unpkg.com/@twind/cdn@1.0.8"), g.Attr("crossorigin", "anonymous")),
				Script(Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js"), g.Attr("crossorigin", "anonymous")),
				Script(Src("https://unpkg.com/idiomorph@0.3.0/dist/idiomorph-ext.min.js"), g.Attr("crossorigin", "anonymous")),
				// Enable hover and active states on touchscreens, https://stackoverflow.com/a/56140328
				Script(g.Raw("document.addEventListener(\"touchstart\", function () {}, true);")),
				Script(g.Raw(htmxScript)),
			),
			Body(Classes(
				Class("grid m-0 min-h-[100vh]"),
				// CSS styles for body element to fill visible area in safari browser on ipad.
				Class("supports-[-webkit-touch-callout:none]:min-h-[-webkit-fill-available]"),
				Class("_overflow-hidden _overscroll-none"),
			),
				TabIndex("0"),
				Div(Class("hidden"),
					Div(
						// Trigger browser refresh after development webserver changes using endpoint polling.
						hx.Get(fmt.Sprintf("/site-refresh?timestamp=%v", serviceTimestamp)),
						hx.Trigger(fmt.Sprintf("every %vms", 250)),
						hx.Swap("none swap:0s settle:0s"),
					),
				),
				Main(Class("relative grid w-full h-full"),
					hx.Ext("morph"),
					hx.Get(fmt.Sprintf("/site?page=%s", pageParam)),
					hx.Trigger("load"),
					Span(g.Text("Loading site...")),
				),
			),
		),
	).Render(writer)
}

func handlerSite(writer http.ResponseWriter, req *http.Request) {
	pageParam := req.URL.Query().Get("page")
	rootNode := router.BuildRootNode(pageParam)
	writer.Header().Set("content-type", "text/html; charset=utf-8")
	writer.Header().Set("hx-push-url", fmt.Sprintf("/pages/%s", pageParam))
	_ = Div(
		Class("grid h-full w-full"),
		g.If(rootNode != nil, Lazy(func() g.Node {
			return rootNode
		})),
	).Render(writer)
}

func handlerSiteRefresh(writer http.ResponseWriter, req *http.Request) {
	clientTimestamp := req.URL.Query().Get("timestamp")

	writer.Header().Set("content-type", "text/plain; charset=utf-8")
	if clientTimestamp != "" && clientTimestamp != serviceTimestamp {
		writer.Header().Set("hx-refresh", "true")
	}
	writer.WriteHeader(http.StatusOK)
}

const htmxScript string = /* javascript */ `
htmx.onLoad(() => {
	const positionLeftBottom = (trigger, container) => {
		const triggerRect = trigger.getBoundingClientRect();
		const offsetLeft = 0;
		const offsetTop = window.scrollY;
		container.style.left = triggerRect.left + offsetLeft + "px";
		container.style.top = triggerRect.bottom + offsetTop + "px";
	};

	const positionCenterBottom = (trigger, container) => {
		const triggerRect = trigger.getBoundingClientRect();
		const containerRect = container.getBoundingClientRect();
		const offsetLeft = -(containerRect.width / 2);
		const offsetTop = window.scrollY + 8;
		container.style.left = triggerRect.left + offsetLeft + "px";
		container.style.top = triggerRect.bottom + offsetTop + "px";
	};

	htmx.showDialog = (event, dialog) => {
		dialog.showModal();
		const button = event.currentTarget;
		if (button) {
			positionLeftBottom(button, dialog);
		}
	};

	htmx.hideDialog = (event, dialog) => {
		const maybeDialog = event.target;
		if (maybeDialog === dialog) {
			dialog.close();
		}
	};

	htmx.showPopover = (event, popover) => {
		// popover.showPopover();
		const button = event.currentTarget;
		const handleToggleEvent = (event) => {
			popover.removeEventListener("toggle", handleToggleEvent);
			if (button) {
				positionCenterBottom(button, popover);
			}
		};
		popover.addEventListener("toggle", handleToggleEvent);
	}

	htmx.hidePopover = (event, popover) => {
		popover.hidePopover();
		const maybePopover = event.target;
		if (maybePopover === popover) {
			popover.hidePopover();
		}
	};
});
`
