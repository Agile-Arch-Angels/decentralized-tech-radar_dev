# Preface

*Rendering of the tech radar was strongly inspired by the open source tech radar project created by Thoughtworks: https://github.com/thoughtworks/build-your-own-radar/*

# Changing quadrant and ring names

To change the existing quadrant and ring names, navigate to ./src/js/graphing/config.js on line 5 and 6. Alter the string names and save.

# Styling

To change the style of the tech radar, navigate to ./src/js/stylesheets/style.css and start editing.

# Adding or removing inner Blip data 

Blip data can be altered in ../src/js/util/inputSanitizer.js (.. to be updated)

# Images

Tech radar specific images are linked from ./src/js/images and utilized through style.css.

## Dependencies
Rendering a tech radar requires certain javascript libraries. These dependencies are all isolated offline and integrated into the project through script injection in the html.

RenderJS is one library, which makes sure multiple js files are able to be used together. RenderJS is integral for the current program to work. Navigate to ./src/HTML/makeHtml.go from line 45. Here the RenderJS library is linked, where the following short scripts utilize the library. The scripts using RequireJS are requireConfig.js (with further library dependencies from other sources) and a script to start the js files which then builds the tech radar. 

For injecting scripts from html side:
```
    <script src="Path to script"></script>
```

Example of injecting d3:
```
    <script src="https://d3js.org/d3.v7.min.js"></script>
```

To add or change library dependencies in requireConfig.js, navigate to ./src/js/requireConfig.js.

For injecting script from the javascript side:
```
    'NameOfLibrary': 'PathToScriptLibrary'
```

Example of injecting d3 (in this case the path is a url link, but could also be a file path):
```
    'd3': 'https://d3js.org/d3.v7.min'
```

If one JS library requires another library, this can be added in requireConfig.js as a "Shim". In the example below, d3tip loads after d3, as d3tip requires d3 to properly work.
```
    shim: {
        'd3tip': {
            deps: ['d3','d3-collection', 'd3-selection'],
            exports: 'd3.tip'
        }, ...
```
For the example, the items in "deps" (short for dependency) has to be in first part of requireConfig.js (shown above)
