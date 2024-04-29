define(['d3'], function(d3) {
  const addPdfCoverTitle = (title) => {
    d3.select('main #pdf-cover-page .pdf-title').text(title)
  }

  const addRadarLinkInPdfView = () => {
    d3.select('#generated-radar-link').attr('href', window.location.href)
  }

  const addQuadrantNameInPdfView = (order, quadrantName) => {
    d3.select(`.quadrant-table.${order}`)
      .insert('div', ':first-child')
      .attr('class', 'quadrant-table__name')
      .text(quadrantName)
  }

  return {
    addPdfCoverTitle: addPdfCoverTitle,
    addQuadrantNameInPdfView: addQuadrantNameInPdfView,
    addRadarLinkInPdfView: addRadarLinkInPdfView
  }
});