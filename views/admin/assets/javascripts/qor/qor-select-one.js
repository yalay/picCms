(function (factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as anonymous module.
    define(['jquery'], factory);
  } else if (typeof exports === 'object') {
    // Node / CommonJS
    factory(require('jquery'));
  } else {
    // Browser globals.
    factory(jQuery);
  }
})(function ($) {

  'use strict';

  var $body = $('body');
  var $document = $(document);
  var Mustache = window.Mustache;
  var NAMESPACE = 'qor.selectone';
  var PARENT_NAMESPACE = 'qor.bottomsheets';
  var EVENT_CLICK = 'click.' + NAMESPACE;
  var EVENT_ENABLE = 'enable.' + NAMESPACE;
  var EVENT_DISABLE = 'disable.' + NAMESPACE;
  var EVENT_RELOAD = 'reload.' + PARENT_NAMESPACE;
  var CLASS_CLEAR_SELECT = '.qor-selected__remove';
  var CLASS_CHANGE_SELECT = '.qor-selected__change';
  var CLASS_SELECT_FIELD = '.qor-field__selected';
  var CLASS_SELECT_INPUT = '.qor-field__selectone-input';
  var CLASS_SELECT_TRIGGER = '.qor-field__selectone-trigger';
  var CLASS_PARENT = '.qor-field__selectone';
  var CLASS_BOTTOMSHEETS = '.qor-bottomsheets';
  var CLASS_SELECTED = 'is_selected';
  var CLASS_ONE = 'qor-bottomsheets__select-one';


  function QorSelectOne(element, options) {
    this.$element = $(element);
    this.options = $.extend({}, QorSelectOne.DEFAULTS, $.isPlainObject(options) && options);
    this.init();
  }

  QorSelectOne.prototype = {
    constructor: QorSelectOne,

    init: function () {
      this.bind();
    },

    bind: function () {
      $document.on(EVENT_CLICK, '[data-selectone-url]', this.openBottomSheets.bind(this)).
                on(EVENT_RELOAD, '.' + CLASS_ONE, this.reloadData.bind(this));

      this.$element.
        on(EVENT_CLICK, CLASS_CLEAR_SELECT, this.clearSelect.bind(this)).
        on(EVENT_CLICK, CLASS_CHANGE_SELECT, this.changeSelect);
    },

    unbind: function () {
      $document.off(EVENT_CLICK, '[data-selectone-url]', this.openBottomSheets.bind(this)).
                off(EVENT_RELOAD, '.' + CLASS_ONE, this.reloadData.bind(this));

      this.$element.
        off(EVENT_CLICK, CLASS_CLEAR_SELECT, this.clearSelect.bind(this)).
        off(EVENT_CLICK, CLASS_CHANGE_SELECT, this.changeSelect);
    },

    clearSelect: function (e) {
      var $target = $(e.target),
          $parent = $target.closest(CLASS_PARENT);

      $parent.find(CLASS_SELECT_FIELD).remove();
      $parent.find(CLASS_SELECT_INPUT)[0].value = '';
      $parent.find(CLASS_SELECT_TRIGGER).show();

      $parent.trigger('qor.selectone.unselected');
      return false;
    },

    changeSelect: function () {
      var $target = $(this),
          $parent = $target.closest(CLASS_PARENT);

      $parent.find(CLASS_SELECT_TRIGGER).trigger('click');

    },

    openBottomSheets: function (e) {
      var $this = $(e.target),
          data = $this.data();

      this.BottomSheets = $body.data('qor.bottomsheets');
      this.bottomsheetsData = data;
      this.$parent = $this.closest(CLASS_PARENT);

      data.url = data.selectoneUrl;

      this.SELECT_ONE_SELECTED_ICON = $('[name="select-one-selected-icon"]').html();

      this.BottomSheets.open(data, this.handleSelectOne.bind(this));
    },

    initItem: function () {
      var $selectFeild = this.$parent.find(CLASS_SELECT_FIELD),
          selectedID;

      if (!$selectFeild.length) {
        return;
      }

      selectedID = $selectFeild.data().primaryKey;

      if (selectedID) {
        $(CLASS_BOTTOMSHEETS).find('tr[data-primary-key="' + selectedID + '"]').addClass(CLASS_SELECTED).find('td:first').append(this.SELECT_ONE_SELECTED_ICON);
      }
    },

    reloadData: function () {
      this.initItem();
    },

    renderSelectOne: function (data) {
      return Mustache.render($('[name="select-one-selected-template"]').html(), data);
    },

    handleSelectOne: function () {
      var options = {
        onSelect: this.onSelectResults.bind(this), //render selected item after click item lists
        onSubmit: this.onSubmitResults.bind(this)  //render new items after new item form submitted
      };

      $(CLASS_BOTTOMSHEETS).qorSelectCore(options).addClass(CLASS_ONE).data(this.bottomsheetsData);
      this.initItem();
    },

    onSelectResults: function (data) {
      this.handleResults(data);
    },

    onSubmitResults: function (data) {
      this.handleResults(data, true);
    },

    handleResults: function (data, isNewData) {
      var template,
          bottomsheetsData = this.bottomsheetsData,
          $parent = this.$parent,
          $select = bottomsheetsData.selectId ? $(bottomsheetsData.selectId) : $parent.find('select'),
          $selectFeild = $parent.find(CLASS_SELECT_FIELD);

      data.displayName = data.Text || data.Name || data.Title || data.Code || data[Object.keys(data)[0]];

      if (!$select.length) {
        return;
      }

      $select[0].value = data.primaryKey;
      template = this.renderSelectOne(data);

      if ($selectFeild.length) {
        $selectFeild.remove();
      }

      $parent.prepend(template);
      $parent.find(CLASS_SELECT_TRIGGER).hide();

      if (isNewData) {
        $select.append(Mustache.render(QorSelectOne.SELECT_ONE_OPTION_TEMPLATE, data));
        $select[0].value = data.primaryKey;
      }

      $parent.trigger('qor.selectone.selected', [data]);

      $(CLASS_BOTTOMSHEETS).qorSelectCore('destroy');
      this.BottomSheets.hide();
    },

    destroy: function () {
      this.unbind();
      this.$element.removeData(NAMESPACE);
    }

  };

  QorSelectOne.SELECT_ONE_OPTION_TEMPLATE = '<option value="[[ primaryKey ]]" >[[ displayName ]]</option>';

  QorSelectOne.plugin = function (options) {
    return this.each(function () {
      var $this = $(this);
      var data = $this.data(NAMESPACE);
      var fn;

      if (!data) {
        if (/destroy/.test(options)) {
          return;
        }

        $this.data(NAMESPACE, (data = new QorSelectOne(this, options)));
      }

      if (typeof options === 'string' && $.isFunction(fn = data[options])) {
        fn.apply(data);
      }
    });
  };

  $(function () {
    var selector = '[data-toggle="qor.selectone"]';
    $(document).
      on(EVENT_DISABLE, function (e) {
        QorSelectOne.plugin.call($(selector, e.target), 'destroy');
      }).
      on(EVENT_ENABLE, function (e) {
        QorSelectOne.plugin.call($(selector, e.target));
      }).
      triggerHandler(EVENT_ENABLE);
  });

  return QorSelectOne;

});
