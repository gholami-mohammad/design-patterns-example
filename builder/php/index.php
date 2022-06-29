<?php

class Page
{
    public string $title;
    public ?string $content;
    public string $type;
}

abstract class PageBuilder
{
    protected Page $page;
    public function __construct(protected string $title, protected ?string $content)
    {
        $this->page = new Page();
        $this->page->type = $this->getType();
    }

    public function getPage(): Page
    {
        return $this->page;
    }

    abstract public function addTitle(): PageBuilder;
    abstract public function addContent(): PageBuilder;
    abstract public function getType(): string;
}

class HTMLPageBuilder extends PageBuilder
{
    public function addTitle(): PageBuilder
    {
        $this->page->title = "<h1>$this->title</h1>";
        return $this;
    }

    public function addContent(): PageBuilder
    {
        $this->page->content = "<div class=\"content\">$this->content</div>";
        return $this;
    }
    
    public function getType(): string
    {
        return 'html';
    }
}

class MarkdownPageBuilder extends PageBuilder
{
    public function addTitle(): PageBuilder
    {
        $this->page->title = "# $this->title";
        return $this;
    }

    public function addContent(): PageBuilder
    {
        $this->page->content = $this->content;
        return $this;
    }
    
    public function getType(): string
    {
        return 'md';
    }
}


class Page404Error extends HTMLPageBuilder
{
    public function addTitle(): PageBuilder
    {
        $this->page->title = "<h1 class=\"page404-title\">$this->title</h1>";
        return $this;
    }

    public function addContent(): PageBuilder
    {
        $this->page->content = "<div class=\"content\">$this->content</div>";
        return $this;
    }
}

class Page500Error extends HTMLPageBuilder
{
    public function addTitle(): PageBuilder
    {
        $this->page->title = "<h1 class=\"page500-title\">$this->title</h1>";
        return $this;
    }

    public function addContent(): PageBuilder
    {
        $this->page->content = "<div class=\"content\">$this->content</div>";
        return $this;
    }
}

class PageCreator
{
    public function buildPage(PageBuilder $builder): Page
    {
        $builder->addTitle();
        $builder->addContent();
        return $builder->getPage();
    }
}

// Usage
$pageCreator = new PageCreator();

$htmlBuilder = new HTMLPageBuilder('Builder Design Pattenr', '<div>This is an exmaple of builder design pattern</div>');
$htmlPage = $pageCreator->buildPage($htmlBuilder);
print_r($htmlPage);
echo '------------------------------------------------------------';

$mdBuilder = new MarkdownPageBuilder('Builder Design Pattenr', 'This is an exmaple of builder design pattern');
$mdPage = $pageCreator->buildPage($mdBuilder);
print_r($mdPage);
echo '------------------------------------------------------------';

$err404Builder = new Page404Error('Nothing is here', null);
$err404 = $pageCreator->buildPage($err404Builder);
print_r($htmlPage);
echo '------------------------------------------------------------';
