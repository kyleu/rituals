package site

import (
	"fmt"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/site/download"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/doc"
	"github.com/kyleu/rituals/views"
	"github.com/kyleu/rituals/views/layout"
	"github.com/kyleu/rituals/views/verror"
	"github.com/kyleu/rituals/views/vsite"
)

func Handle(as *app.State, path []string, ps *cutil.PageState) (string, layout.Page, []string, error) {
	ps.SearchPath = "-"
	if len(path) == 0 {
		ps.Data = siteData("Welcome to the marketing site!")
		return "", &vsite.Index{}, path, nil
	}

	var page layout.Page
	var err error
	switch path[0] {
	case util.AppKey:
		msg := "\n  " +
			"<meta name=\"go-import\" content=\"github.com/kyleu/rituals git %s\">\n  " +
			"<meta name=\"go-source\" content=\"github.com/kyleu/rituals %s %s/tree/main{/dir} %s/blob/main{/dir}/{file}#L{line}\">"
		ps.HeaderContent = fmt.Sprintf(msg, util.AppSource, util.AppSource, util.AppSource, util.AppSource)
		return "", &vsite.GoSource{}, []string{"Source**code"}, nil
	case keyAbout:
		ps.SetTitleAndData("About "+util.AppName, util.AppName+" v"+as.BuildInfo.Version)
		page = &views.About{Version: as.BuildInfo.Version, Started: as.Started}
	case keyDownload:
		dls := download.GetLinks(as.BuildInfo.Version)
		data := util.ValueMap{"base": "https://github.com/kyleu/rituals/releases/download/v" + as.BuildInfo.Version, "links": dls}
		ps.SetTitleAndData("Downloads", data)
		page = &vsite.Download{Links: dls, Version: as.BuildInfo.Version}
	case keyInstall:
		page, err = mdTemplate("This static page contains installation instructions", "installation.md", "code", ps)
	case keyContrib:
		page, err = mdTemplate("This static page describes how to build "+util.AppName, "contributing.md", "cog", ps)
	case keyTech:
		page, err = mdTemplate("This static page describes the technology used in "+util.AppName, "technology.md", "shield", ps)
	default:
		page, err = mdTemplate("Documentation for "+util.AppName, path[0]+util.ExtMarkdown, "", ps)
		if err != nil {
			page = &verror.NotFound{Path: "/" + util.StringJoin(path, "/")}
			err = nil
		}
	}
	return "", page, path, err
}

func siteData(result string, kvs ...string) util.ValueMap {
	ret := util.ValueMap{"app": util.AppName, "url": util.AppURL, "result": result}
	for i := 0; i < len(kvs); i += 2 {
		ret[kvs[i]] = kvs[i+1]
	}
	return ret
}

func mdTemplate(description string, path string, icon string, ps *cutil.PageState) (layout.Page, error) {
	if icon == "" {
		icon = "cog"
	}
	title, html, err := doc.HTML(path, path, func(s string) (string, string, error) {
		return cutil.FormatMarkdownClean(s, icon)
	})
	if err != nil {
		return nil, err
	}
	ps.SetTitleAndData(title, siteData(title, "description", description))
	page := &vsite.MarkdownPage{Title: title, HTML: html}
	return page, nil
}
