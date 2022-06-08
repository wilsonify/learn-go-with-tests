HTML Templates
========

We live in a world where everyone wants to build web applications
with the latest flavour of the month frontend framework built upon
gigabytes of transpiled JavaScript, working with a Byzantine build system; but maybe that’s not always necessary.

I’d say most Go developers value a simple, stable & fast toolchain but
the frontend world frequently fails to deliver on this front.

Many websites do not need to be an SPA. 

HTML and CSS are fantastic ways of delivering content and you can use Go to make a
website to deliver HTML.

If you wish to still have some dynamic elements, 
you can still sprinkle in some client side JavaScript, 
or you may even want to try experimenting with Hotwire which allows you to deliver a dynamic experience with a server-side approach.

You can generate your HTML in Go with elaborate usage of fmt.Fprintf,
but in this chapter you’ll learn that Go’s standard library has some
tools to generate HTML in a simpler and more maintainable way. 

You’ll also learn more effective ways of testing this kind of code that you
may not have run in to before.

# What we’re going to build
In the Reading Files chapter we wrote some code that would take an fs.FS (a file-system), 
and return a slice of Post for each markdown file it encountered.

Here’s an example of one of the markdown files that can be parsed.
```
Title: Welcome to my blog
Description: Introduction to my blog
Tags: cooking, family, live-laugh-love
---
# First recipe!
Welcome to my **amazing recipe blog**. I am going to write about my family recipes, and make
```

If we continue our journey of writing blog software, we’d take this data
and generate HTML from it for our web server to return in response
to HTTP requests.

For our blog, we want to generate two kinds of page:

1. View post. Renders a specific post. The Body field in Post is a string containing markdown so that should be converted to HTML.

2. Index. Lists all of the posts, with hyperlinks to view the specific post.

We’ll also want a consistent look and feel across our site, 
so for each page we’ll have the usual HTML furniture like <html> and a <head>
containing links to CSS stylesheets and whatever else we may want.

When you’re building blog software you have a few options in terms
of approach of how you build and send HTML to the user’s browser.

We’ll design our code so it accepts an io.Writer. 

This means the caller of our code has the flexibility to:

* Write them to an os.File , so they can be statically served

* Write out the HTML directly to a http.ResponseWriter

* Or just write them to anything really! So long as it implements

io.Writer the user can generate some HTML from a Post

# Write the test first

As always, 
it’s important to think about requirements before diving in too fast. 

How can we take this large-ish set of requirements and break it down in to a small, 
achievable step that we can focus on?

In my view, actually viewing content is higher priority than an index
page. 

We could launch this product and share direct links to our wonderful content. 
An index page which cant link to the actual content isn’t useful.

Still, rendering a post as described earlier still feels big. 

All the HTML furniture, 
converting the body markdown into HTML, 
listing tags, 
e.t.c.
At this stage I’m not overly concerned with the specific markup, 
and an easy first step would be just to check we can render the post’s title
as an <h1>. 

This feels like the smallest first step that can move us forward a bit.


# Introducing templates

Go has two templating packages text/template and html/template
and they share the same interface. 

What they both do is allow you to combine a template and some data to produce a string.

What’s the difference with the HTML version? 

Package template (html/template) implements data-driven templates for generating HTML output safe against code injection.

It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.

The templating language is very similar to Mustache and allows you to dynamically generate content in a very clean fashion with a nice separation of concerns. 

Compared to other templating languages you may have used, 
it is very constrained or ”logic-less” as Mustache likes to say. 

This is an important, and deliberate design decision. 

Whilst we’re focusing on generating HTML here, 
if your project is doing complex string concatenations and incantations, 
you might want to reach for text/template to clean up your code.


We create a new template with a name, 
and then parse our template string. 

We can then use the Execute method on it, passing in our data, 
in this case the Post.

The template will substitute things like {{.Description}} with the content of 
p.Description. 

Templates also give you some programming primitives like range to loop over values, 
and if. 

You can find more details in the text/template documentation.

Importantly, our code is easier to read and has far less annoying error handling to contend with.

Frequently people complain about the verbosity of error handling in Go, 
but you might find you can find better ways to write your code so
it’s less error-prone in the first place, like here.

