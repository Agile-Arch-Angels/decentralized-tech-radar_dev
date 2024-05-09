// https://d3js.org/d3-collection/ v1.0.7 Copyright 2018 Mike Bostock
!(function (n, t) {
  "object" == typeof exports && "undefined" != typeof module
    ? t(exports)
    : "function" == typeof define && define.amd
    ? define(["exports"], t)
    : t((n.d3 = n.d3 || {}));
})(this, function (n) {
  "use strict";
  function t() {}
  function e(n, e) {
    var r = new t();
    if (n instanceof t)
      n.each(function (n, t) {
        r.set(t, n);
      });
    else if (Array.isArray(n)) {
      var i,
        u = -1,
        o = n.length;
      if (null == e) for (; ++u < o; ) r.set(u, n[u]);
      else for (; ++u < o; ) r.set(e((i = n[u]), u, n), i);
    } else if (n) for (var s in n) r.set(s, n[s]);
    return r;
  }
  function r() {
    return {};
  }
  function i(n, t, e) {
    n[t] = e;
  }
  function u() {
    return e();
  }
  function o(n, t, e) {
    n.set(t, e);
  }
  function s() {}
  t.prototype = e.prototype = {
    constructor: t,
    has: function (n) {
      return "$" + n in this;
    },
    get: function (n) {
      return this["$" + n];
    },
    set: function (n, t) {
      return (this["$" + n] = t), this;
    },
    remove: function (n) {
      var t = "$" + n;
      return t in this && delete this[t];
    },
    clear: function () {
      for (var n in this) "$" === n[0] && delete this[n];
    },
    keys: function () {
      var n = [];
      for (var t in this) "$" === t[0] && n.push(t.slice(1));
      return n;
    },
    values: function () {
      var n = [];
      for (var t in this) "$" === t[0] && n.push(this[t]);
      return n;
    },
    entries: function () {
      var n = [];
      for (var t in this)
        "$" === t[0] && n.push({ key: t.slice(1), value: this[t] });
      return n;
    },
    size: function () {
      var n = 0;
      for (var t in this) "$" === t[0] && ++n;
      return n;
    },
    empty: function () {
      for (var n in this) if ("$" === n[0]) return !1;
      return !0;
    },
    each: function (n) {
      for (var t in this) "$" === t[0] && n(this[t], t.slice(1), this);
    },
  };
  var f = e.prototype;
  function c(n, t) {
    var e = new s();
    if (n instanceof s)
      n.each(function (n) {
        e.add(n);
      });
    else if (n) {
      var r = -1,
        i = n.length;
      if (null == t) for (; ++r < i; ) e.add(n[r]);
      else for (; ++r < i; ) e.add(t(n[r], r, n));
    }
    return e;
  }
  (s.prototype = c.prototype =
    {
      constructor: s,
      has: f.has,
      add: function (n) {
        return (this["$" + (n += "")] = n), this;
      },
      remove: f.remove,
      clear: f.clear,
      values: f.keys,
      size: f.size,
      empty: f.empty,
      each: f.each,
    }),
    (n.nest = function () {
      var n,
        t,
        s,
        f = [],
        c = [];
      function a(r, i, u, o) {
        if (i >= f.length) return null != n && r.sort(n), null != t ? t(r) : r;
        for (
          var s, c, h, l = -1, v = r.length, p = f[i++], y = e(), d = u();
          ++l < v;

        )
          (h = y.get((s = p((c = r[l])) + ""))) ? h.push(c) : y.set(s, [c]);
        return (
          y.each(function (n, t) {
            o(d, t, a(n, i, u, o));
          }),
          d
        );
      }
      return (s = {
        object: function (n) {
          return a(n, 0, r, i);
        },
        map: function (n) {
          return a(n, 0, u, o);
        },
        entries: function (n) {
          return (function n(e, r) {
            if (++r > f.length) return e;
            var i,
              u = c[r - 1];
            return (
              null != t && r >= f.length
                ? (i = e.entries())
                : ((i = []),
                  e.each(function (t, e) {
                    i.push({ key: e, values: n(t, r) });
                  })),
              null != u
                ? i.sort(function (n, t) {
                    return u(n.key, t.key);
                  })
                : i
            );
          })(a(n, 0, u, o), 0);
        },
        key: function (n) {
          return f.push(n), s;
        },
        sortKeys: function (n) {
          return (c[f.length - 1] = n), s;
        },
        sortValues: function (t) {
          return (n = t), s;
        },
        rollup: function (n) {
          return (t = n), s;
        },
      });
    }),
    (n.set = c),
    (n.map = e),
    (n.keys = function (n) {
      var t = [];
      for (var e in n) t.push(e);
      return t;
    }),
    (n.values = function (n) {
      var t = [];
      for (var e in n) t.push(n[e]);
      return t;
    }),
    (n.entries = function (n) {
      var t = [];
      for (var e in n) t.push({ key: e, value: n[e] });
      return t;
    }),
    Object.defineProperty(n, "__esModule", { value: !0 });
});
