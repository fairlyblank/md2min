package main

var templContent = `<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="chrome=1">
    <title>Fairlyblank.github.com by fairlyblank</title>

    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Lato:300italic,700italic,300,700">		
		<style type="text/css">
body {
  padding:50px;
  font:14px/1.5 Lato, "Helvetica Neue", Helvetica, Arial, sans-serif;
  color:#777;
  font-weight:300;
}
h1, h2, h3, h4, h5, h6 {
  color:#222;
  margin:0 0 20px;
}
p, ul, ol, table, pre, dl {
  margin:0 0 20px;
}
h1, h2, h3 {
  line-height:1.1;
}
h1 {
  font-size:28px;
}
h2 {
  color:#393939;
}
h3, h4, h5, h6 {
  color:#494949;
}
a {
  color:#39c;
  font-weight:400;
  text-decoration:none;
}
a small {
  font-size:11px;
  color:#777;
  margin-top:-0.6em;
  display:block;
}
.wrapper {
  width:860px;
  margin:0 auto;
}
blockquote {
  border-left:1px solid #e5e5e5;
  margin:0;
  padding:0 0 0 20px;
  font-style:italic;
}
code, pre {
  font-family:Monaco, Bitstream Vera Sans Mono, Lucida Console, Terminal;
  color:#333;
  font-size:12px;
}
pre {
  padding:8px 15px;
  background: #f8f8f8;  
  border-radius:5px;
  border:1px solid #e5e5e5;
  overflow-x: auto;
}
table {
  width:100%;
  border-collapse:collapse;
}
th, td {
  text-align:left;
  padding:5px 10px;
  border-bottom:1px solid #e5e5e5;
}
dt {
  color:#444;
  font-weight:700;
}
th {
  color:#444;
}
img {
  max-width:100%;
}
header {
  width:270px;
  float:left;
  position:fixed;
}
header div {
	display:block;
	overflow-x:hidden;
	overflow-y:scroll;
	height:400px;
}
::-webkit-scrollbar {
	width: 4px;
	height: 8px;
}

::-webkit-scrollbar-track-piece {
	background-color: #f0f0f0;
	border-radius: 4px;
}

::-webkit-scrollbar-thumb {
	background-color: #cfcfcf;
	border-radius: 4px;
}
header ul {
	margin:0;
	padding:0;
  list-style:none;
  height:40px;
  padding:0;
  width:270px;
}
header li {
	margin:0;
	padding:2px;
#  border-right:1px solid #d2d2d2;
  height:20px;
}
header ul a {
  line-height:1;
  text-align:center;
  padding-top:0;
  height:40px;
}
strong {
  color:#222;
  font-weight:700;
}
header ul a strong {
  font-size:14px;
	display:block;
  color:#222;
}
section {
  width:500px;
  float:right;
  padding-bottom:50px;
}
small {
  font-size:11px;
}
hr {
  border:0;
  background:#e5e5e5;
  height:1px;
  margin:0 0 20px;
}
footer {
  width:270px;
  float:left;
  position:fixed;
  bottom:50px;
}
@media print, screen and (max-width: 960px) {
  div.wrapper {
    width:auto;
    margin:0;
  }
  header, section, footer {
    float:none;
    position:static;
    width:auto;
  }
  header {
    padding-right:320px;
  }
  section {
    border:1px solid #e5e5e5;
    border-width:1px 0;
    padding:20px 0;
    margin:0 0 20px;
  }
  header a small {
    display:inline;
  }
  header ul {
    position:absolute;
    right:50px;
    top:52px;
  }
}
@media print, screen and (max-width: 720px) {
  body {
    word-wrap:break-word;
  }
  header {
    padding:0;
  }
  header ul, header p.view {
    position:static;
  }
  pre, code {
    word-wrap:normal;
  }
}
@media print, screen and (max-width: 480px) {
  body {
    padding:15px;
  }
  header ul {
    display:none;
  }
}
@media print {
  body {
    padding:0.4in;
    font-size:12pt;
    color:#444;
  }
}

.highlight  { background: #ffffff; }
.highlight .c { color: #999988; font-style: italic } /* Comment */
.highlight .err { color: #a61717; background-color: #e3d2d2 } /* Error */
.highlight .k { font-weight: bold } /* Keyword */
.highlight .o { font-weight: bold } /* Operator */
.highlight .cm { color: #999988; font-style: italic } /* Comment.Multiline */
.highlight .cp { color: #999999; font-weight: bold } /* Comment.Preproc */
.highlight .c1 { color: #999988; font-style: italic } /* Comment.Single */
.highlight .cs { color: #999999; font-weight: bold; font-style: italic } /* Comment.Special */
.highlight .gd { color: #000000; background-color: #ffdddd } /* Generic.Deleted */
.highlight .gd .x { color: #000000; background-color: #ffaaaa } /* Generic.Deleted.Specific */
.highlight .ge { font-style: italic } /* Generic.Emph */
.highlight .gr { color: #aa0000 } /* Generic.Error */
.highlight .gh { color: #999999 } /* Generic.Heading */
.highlight .gi { color: #000000; background-color: #ddffdd } /* Generic.Inserted */
.highlight .gi .x { color: #000000; background-color: #aaffaa } /* Generic.Inserted.Specific */
.highlight .go { color: #888888 } /* Generic.Output */
.highlight .gp { color: #555555 } /* Generic.Prompt */
.highlight .gs { font-weight: bold } /* Generic.Strong */
.highlight .gu { color: #800080; font-weight: bold; } /* Generic.Subheading */
.highlight .gt { color: #aa0000 } /* Generic.Traceback */
.highlight .kc { font-weight: bold } /* Keyword.Constant */
.highlight .kd { font-weight: bold } /* Keyword.Declaration */
.highlight .kn { font-weight: bold } /* Keyword.Namespace */
.highlight .kp { font-weight: bold } /* Keyword.Pseudo */
.highlight .kr { font-weight: bold } /* Keyword.Reserved */
.highlight .kt { color: #445588; font-weight: bold } /* Keyword.Type */
.highlight .m { color: #009999 } /* Literal.Number */
.highlight .s { color: #d14 } /* Literal.String */
.highlight .na { color: #008080 } /* Name.Attribute */
.highlight .nb { color: #0086B3 } /* Name.Builtin */
.highlight .nc { color: #445588; font-weight: bold } /* Name.Class */
.highlight .no { color: #008080 } /* Name.Constant */
.highlight .ni { color: #800080 } /* Name.Entity */
.highlight .ne { color: #990000; font-weight: bold } /* Name.Exception */
.highlight .nf { color: #990000; font-weight: bold } /* Name.Function */
.highlight .nn { color: #555555 } /* Name.Namespace */
.highlight .nt { color: #000080 } /* Name.Tag */
.highlight .nv { color: #008080 } /* Name.Variable */
.highlight .ow { font-weight: bold } /* Operator.Word */
.highlight .w { color: #bbbbbb } /* Text.Whitespace */
.highlight .mf { color: #009999 } /* Literal.Number.Float */
.highlight .mh { color: #009999 } /* Literal.Number.Hex */
.highlight .mi { color: #009999 } /* Literal.Number.Integer */
.highlight .mo { color: #009999 } /* Literal.Number.Oct */
.highlight .sb { color: #d14 } /* Literal.String.Backtick */
.highlight .sc { color: #d14 } /* Literal.String.Char */
.highlight .sd { color: #d14 } /* Literal.String.Doc */
.highlight .s2 { color: #d14 } /* Literal.String.Double */
.highlight .se { color: #d14 } /* Literal.String.Escape */
.highlight .sh { color: #d14 } /* Literal.String.Heredoc */
.highlight .si { color: #d14 } /* Literal.String.Interpol */
.highlight .sx { color: #d14 } /* Literal.String.Other */
.highlight .sr { color: #009926 } /* Literal.String.Regex */
.highlight .s1 { color: #d14 } /* Literal.String.Single */
.highlight .ss { color: #990073 } /* Literal.String.Symbol */
.highlight .bp { color: #999999 } /* Name.Builtin.Pseudo */
.highlight .vc { color: #008080 } /* Name.Variable.Class */
.highlight .vg { color: #008080 } /* Name.Variable.Global */
.highlight .vi { color: #008080 } /* Name.Variable.Instance */
.highlight .il { color: #009999 } /* Literal.Number.Integer.Long */

.type-csharp .highlight .k { color: #0000FF }
.type-csharp .highlight .kt { color: #0000FF }
.type-csharp .highlight .nf { color: #000000; font-weight: normal }
.type-csharp .highlight .nc { color: #2B91AF }
.type-csharp .highlight .nn { color: #000000 }
.type-csharp .highlight .s { color: #A31515 }
.type-csharp .highlight .sc { color: #A31515 }
		</style>
    <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
    <!--[if lt IE 9]>
    <script src="//html5shiv.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
  </head>
  <body>
    <div class="wrapper">
      <header>
        <h1>{{.Title1}}</h1>
        <p class="view">{{.Title2}}</a></p>
				<div>{{.ListMenu}}</div>
				<footer>
					<p><small>Hosted on GitHub Pages &mdash; Theme by <a href="https://github.com/orderedlist">orderedlist</a></small></p>
				</footer>
      </header>
      <section>{{.Content}}</section>
    </div>
    <script type="script/javascript">
var metas = document.getElementsByTagName('meta');
var i;
if (navigator.userAgent.match(/iPhone/i)) {
  for (i=0; i<metas.length; i++) {
    if (metas[i].name == "viewport") {
      metas[i].content = "width=device-width, minimum-scale=1.0, maximum-scale=1.0";
    }
  }
  document.addEventListener("gesturestart", gestureStart, false);
}
function gestureStart() {
  for (i=0; i<metas.length; i++) {
    if (metas[i].name == "viewport") {
      metas[i].content = "width=device-width, minimum-scale=0.25, maximum-scale=1.6";
    }
  }
}
		</script>
  </body>
</html>
`
