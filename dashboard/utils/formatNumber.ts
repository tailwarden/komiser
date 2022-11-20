function formatNumber(number: number, display?: 'full') {
  if (display === 'full') {
    return new Intl.NumberFormat(undefined, {
      notation: 'standard'
    }).format(Number(number));
  }

  return new Intl.NumberFormat(undefined, {
    notation: 'compact',
    compactDisplay: 'short'
  }).format(Number(number));
}

export default formatNumber;
