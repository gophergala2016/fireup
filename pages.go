package main

import (
	"html/template"
)

var RootTemplate = template.Must(template.New("root").Parse(RootTemplateStr))

const RootTemplateStr = `
<html lang="en"><head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="icon" href="/favicon.ico">

    <title>FireUp</title>
    <link href="/assets/css/bootstrap.min.css" rel="stylesheet">
    <link href="/assets/css/fireup.css" rel="stylesheet">
  </head>

  <body>

    <div class="site-wrapper">
      <div class="site-wrapper-inner">
        <div class="cover-container">
          <div class="masthead clearfix"> </div>
          <div class="inner cover">
            <h1 class="cover-heading"><span class="glyphicon glyphicon-fire" aria-hidden="true"></span></h1>
            <p class="lead">No content to display.</p>
						<p class="lead">Push some content with <code>fireup push</code> to display the content here.</p>
          </div>
        </div>
      </div>
    </div>
    <script src="/assets/js/jquery.min.js"></script>
    <script src="/assets/js/bootstrap.min.js"></script>
    <script src="/assets/js/fireup.js"></script>
</body>
</html>
`
