// const d3 = require('d3')
/*
define([
  '../../util/autoComplete.js',
  './quadrants.js'],
 function(AutoComplete, selectRadarQuadrantremoveScrollListener) {
  // const AutoComplete = require('../../util/autoComplete')
  // const { selectRadarQuadrant, removeScrollListener } = require('../components/quadrants')
  const {selectRadarQuadrant, removeScrollListener } = selectRadarQuadrantremoveScrollListener*/

  define([
    '../../util/autoComplete',
    '../components/quadrants',
    'd3'
  ], function(autoCompleteModule, quadrantsModule, d3) {
    const AutoComplete = autoCompleteModule;
    const { selectRadarQuadrant, removeScrollListener } = quadrantsModule;

  function renderSearch(radarHeader, quadrants) {
    const searchContainer = radarHeader.append('div').classed('search-container', true)

    searchContainer
      .append('input')
      .classed('search-container__input', true)
      .attr('placeholder', 'Search this radar')
      .attr('id', 'auto-complete')

    AutoComplete('#auto-complete', quadrants, function (e, ui) {
      const blipId = ui.item.blip.id()
      const quadrant = ui.item.quadrant

      selectRadarQuadrant(quadrant.order, quadrant.startAngle, quadrant.quadrant.name())
      const blipElement = d3.select(
        `.blip-list__item-container[data-blip-id="${blipId}"] .blip-list__item-container__name`,
      )

      removeScrollListener()
      blipElement.dispatch('search-result-click')

      setTimeout(() => {
        blipElement.node().scrollIntoView({
          behavior: 'smooth',
        })
      }, 1500)
    })
  }

  return renderSearch
})
