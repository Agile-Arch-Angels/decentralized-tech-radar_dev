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