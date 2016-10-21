# Latexapi: compiles latex documents via an api call

This is a very simple go app, that accepts a GET/POST call with the param text,
where text is the latex document. To use you need to have xetex installed.

It also contains a Dockerfile, so you can generate nice-looking pdfs everywhere.

## Build

    docker build -t swerter/latexapi_build -f Dockerfile.build .
    docker run --rm -v "$PWD":/usr/src/myapp -it swerter/latexapi_build
    docker build -t swerter/latexapi .

## Run

    docker run --name latexapi --rm -p 8080:8080 -it swerter/latexapi

and open http://localhost:8080

You can try it out by entering the following into the textarea:

    \documentclass{article}

    \begin{document}

    \title{Introduction to \LaTeX{}}
    \author{Author's Name}

    \maketitle

    \begin{abstract}
    The abstract text goes here.
    \end{abstract}

    \section{Introduction}
    Here is the text of your introduction.

    \begin{equation}
        \label{simple_equation}
        \alpha = \sqrt{ \beta }
    \end{equation}

    \subsection{Subsection Heading Here}
    Subsection 1

    \section{Conclusion}
    And the conclusion is:

    \end{document}

Prefer the command line?

    curl -X GET http://localhost:8080/compile\?text\=%5Cdocumentclass%7Barticle%7D%0A%0A%5Cbegin%7Bdocument%7D%0A%0A%5Ctitle%7BIntroduction+to+%5CLaTeX%7B%7D%7D%0A%5Cauthor%7BAuthor%27s+Name%7D%0A%0A%5Cmaketitle%0A%0A%5Cbegin%7Babstract%7D%0AThe+abstract+text+goes+here.%0A%5Cend%7Babstract%7D%0A%0A%5Csection%7BIntroduction%7D%0AHere+is+the+text+of+your+introduction.%0A%0A%5Cbegin%7Bequation%7D%0A++++%5Clabel%7Bsimple_equation%7D%0A++++%5Calpha+%3D+%5Csqrt%7B+%5Cbeta+%7D%0A%5Cend%7Bequation%7D%0A%0A%5Csubsection%7BSubsection+Heading+Here%7D%0ASubsection+1%0A%0A%5Csection%7BConclusion%7D%0AAnd+the+conclusion+is%3A%0A%0A%5Cend%7Bdocument%7D -o /tmp/test.pdf

and then open `/tmp/test.pdf`.
