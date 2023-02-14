function formatNumber(number: number, display?: 'full') {
  if (display === 'full') {
    return new Intl.NumberFormat(undefined, {
      notation: 'standard'
    }).format(number);
  }

  return new Intl.NumberFormat(undefined, {
    notation: 'compact',
    compactDisplay: 'short'
  }).format(number);
}

export default formatNumber;
