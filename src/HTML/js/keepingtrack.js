(function() {
    // d3 
  })();
  
  (function() {
    // d3-collection
  })();
  
  (function() {
    // d3-selection
  })();
  
  (function() {
    // d3-tip
    window.d3.tip = d3.tip;
  })();

  (function(chance) {
    // chance
    window.chance = chance;
  })();

  (function(_) {
    // lodash
    window._ = _;
  })();

  (function($) {
    // jquery
    window.$ = $;
  })();

  (function() {
    // jquery ui
    window.AutoComplete = AutoComplete;
  })();
  
  (function() {
    // custom code
  })(window.d3, window.d3col, window.d3sel, window.d3tip, window.chance, window._, window.$, window.AutoComplete);



  /// 2

  (function() {
    // d3 
  })();
  
  (function() {
    // d3-collection
  })();
  
  (function() {
    // d3-selection
  })();
  
  (function() {
    // d3-tip
  })();

  (function() {
    // chance
  })();

  (function() {
    // lodash
  })();

  (function() {
    // jquery
  })();

  (function() {
    // jquery ui
  })();
  
  (function() {
    // custom code
  })();



  /// 3 trying

  (function(global) {
    // Load d3
    global.d3 = d3;
  })(window);
  
  (function(global) {
    // Load d3-collection
    var d3 = global.d3; 
    
  })(window);
  
  (function(global) {
    // Load d3-selection
    var d3 = global.d3;
  })(window);
  
  (function(global) {
    // Load d3-tip
    var d3 = global.d3 || {};
    //code 
    global.d3.tip = d3tip;
  })(window);

  //try 4

  (function() {
    // Define jQuery module
    var JQ = (function() {
        // jQuery code here
        return jQuery; // Return jQuery object
    })();

    // Define jQuery UI (jquery-autocomplete) module
    var AutoComplete = (function() {
        // jQuery UI code here
        return jQuery.ui; // Return jQuery UI object
    })();

    // Define facModel module
    var facModel = (function(JQ, AutoComplete) {
        // Your facModel code here
        function init() {
            // Code using JQ and AutoComplete
        }

        // Expose init function (or any other public functions) as the interface
        return {
            init: init
        };
    })(JQ, AutoComplete);

    // Initialize facModel
    facModel.init();
})();



//try 5
// use $.$, for instance when calling jquery different places

//try 6
(function(global) {
  // Include lodash with no conflicts, if necessary
  // It can also be just included normally if it doesn't conflict with other libraries
  (function() {
    // lodash
  })();

  // Include jQuery
  (function() {
    // jquery
  })(global);

  // Include jQuery UI, which depends on jQuery
  (function($) {
    // jquery ui library
  })(global.jQuery);

  // Include d3 and its related libraries
  (function() {
    // d3 library
    global.d3 = global.d3 || {};
  })();
  
  (function(d3) {
    // d3-collection
    // Ensure it attaches itself to `d3`
  })(global.d3);

  (function(d3) {
    // d3-selection
    // Ensure it attaches itself to `d3`
  })(global.d3);
  
  (function(d3) {
    // d3-tip library, assuming it attaches itself to `d3`
    // d3-tip is dependent on d3, d3-collection, and d3-selection
  })(global.d3);

  // Include chance
  (function() {
    // chance
  })();

  // Your custom script, where you can use all the libraries
  (function($, d3, Chance, _) { 
    // custom code dependent on d3, d3-tip, chance, lodash, jquery, and jquery ui
  })(global.jQuery, global.d3, global.Chance, global._);
  
})(window);


//

    // const sheetData = {{.CSV}} // CSV data here
  	// const sheet = Factory(sheetData);  sheet.build();


    //try 8

    (function(global) {
      // lodash library
      var _ = {
          // imagine these are lodash functions
          chunk: function() { /* ... */ },
          compact: function() { /* ... */ },
          // ... etc.
      };
      global._ = _; // make lodash available globally under the "_" variable name
    })();
    
    //...
    
    (function($, d3, Chance, _) { 
        // custom code dependent on d3, d3-tip, chance, lodash, jquery, and jquery ui
        // Now you can use lodash functions like _.chunk, _.compact, etc.
    })(window.jQuery, window.d3, window.Chance, window._);