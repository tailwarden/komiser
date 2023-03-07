import formatNumber from './formatNumber';

describe('formatNumber outputs', () => {
  test('should format number (over a thousand) in short notation', () => {
    const result = formatNumber(12345);
    expect(result).toBe('12K');
  });

  test('should format number (over a million) in short notation', () => {
    const result = formatNumber(1234567);
    expect(result).toBe('1.2M');
  });

  test('should format number (over a billion) in short notation', () => {
    const result = formatNumber(1234567890);
    expect(result).toBe('1.2B');
  });

  test('should format number (as currency - dollar) in short notation', () => {
    const result = formatNumber(12345, 'currency');
    expect(result).toBe('$12K');
  });

  test('should format number in full notation', () => {
    const result = formatNumber(12345, 'full');
    expect(result).toBe('12,345');
  });

  test('should format number with decimals in full notation', () => {
    const result = formatNumber(12345.67, 'full');
    expect(result).toBe('12,345.67');
  });
});
