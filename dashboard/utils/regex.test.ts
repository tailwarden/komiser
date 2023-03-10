import regex, { required } from './regex';

describe('regex outputs', () => {
  test('should return the required regex', () => {
    const result = required;
    expect(result).toStrictEqual(/./);
  });

  test('should return the email regex', () => {
    const result = regex.email;
    expect(result).toStrictEqual(/^[\w-\\.]+@([\w-]+\.)+[\w-]{2,4}$/);
  });

  test('should return the password regex', () => {
    const result = regex.password;
    expect(result).toStrictEqual(/^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/);
  });

  test('should return the number regex', () => {
    const result = regex.number;
    expect(result).toStrictEqual(/ \d+ /);
  });
});
