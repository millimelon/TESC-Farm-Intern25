all: build clean

build:
	pandoc -V geometry:margin=1in ./*.md -o farm-document.pdf
	pdfunite farm-document.pdf Farm-Diagram.pdf Network-Diagram.pdf Data-Diagram.pdf Server-Diagram.pdf Power-Diagram.pdf farm-project.pdf

clean:
	rm farm-document.pdf
