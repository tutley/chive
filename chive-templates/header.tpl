{{define "header"}}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>{{.Title}}</title>
    <link rel="icon" type="image/png" sizes="32x32" href="/dist/static/img/icons/favicon-32x32.png">
    <link rel="icon" type="image/png" sizes="16x16" href="/dist/static/img/icons/favicon-16x16.png">
    <!--[if IE]><link rel="shortcut icon" href="/static/img/icons/favicon.ico"><![endif]-->
    <!-- Add to home screen for Android and modern mobile browsers -->
    <link rel="manifest" href="/dist/static/manifest.json">
    <meta name="theme-color" content="#4DBA87">
    <!-- Social Metas -->
    {{if .SiteName}}<meta property="og:site_name" content="{{.SiteName}}">{{end}}
    {{if .Title}}<meta property="og:title" content="{{.Title}}">{{end}}
    {{if .Description}}<meta property="og:description" content="{{.Description}}">{{end}}
    {{if .ImageURL}}<meta property="og:image" content="{{.ImageURL}}">{{end}}
    {{if .URL}}<meta property="og:url" content="{{.URL}}">{{end}}
    {{if .ImageURL}}<meta name="twitter:card" content="summary_large_image">{{end}}
    {{if .TwitterUsername}}<meta name="twitter:site" content="{{.TwitterUsername}}">{{end}}
    {{if .Type}}<meta name="og:type" content="{{.Type}}">{{end}}
    <!-- Add to home screen for Safari on iOS -->
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-mobile-web-app-title" content="{{.Title}}">
    <link rel="apple-touch-icon" href="/dist/static/img/icons/apple-touch-icon-152x152.png">
    <!-- Add to home screen for Windows -->
    <meta name="msapplication-TileImage" content="/dist/static/img/icons/msapplication-icon-144x144.png">
    <meta name="msapplication-TileColor" content="#000000">
{{end}}