function formatNumber(number: number, display?: 'full' | 'currency') {
  if (display === 'full') {
    return new Intl.NumberFormat(undefined, {
      notation: 'standard'
    }).format(number);
  }

  if (display === 'currency') {
    return new Intl.NumberFormat('en-US', {
      notation: 'compact',
      compactDisplay: 'short',
      style: 'currency',
      currency: 'USD'
    }).format(number);
  }

  return new Intl.NumberFormat(undefined, {
    notation: 'compact',
    compactDisplay: 'short'
  }).format(number);
}

export default formatNumber;
