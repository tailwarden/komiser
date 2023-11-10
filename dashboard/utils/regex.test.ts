import regex, { required } from './regex';

describe('regex util', () => {
  it('should return the required regex', () => {
    const result = required;
    expect(result).toStrictEqual(/./);
  });

  it('should return the email regex', () => {
    const result = regex.email;
    expect(result).toStrictEqual(/^[\w-\\.]+@([\w-]+\.)+[\w-]{2,4}$/);
  });

  it('should return the password regex', () => {
    const result = regex.password;
    expect(result).toStrictEqual(/^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/);
  });

  it('should return the number regex', () => {
    const result = regex.number;
    expect(result).toStrictEqual(/ \d+ /);
  });
});
