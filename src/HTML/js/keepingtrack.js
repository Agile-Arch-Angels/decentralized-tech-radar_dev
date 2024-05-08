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