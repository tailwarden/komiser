function formatNumber(number: number) {
  return new Intl.NumberFormat('en-US', {
    notation: 'compact',
    compactDisplay: 'short'
  }).format(Number(number));
}

export default formatNumber;
