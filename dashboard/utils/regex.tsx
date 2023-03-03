type Regex = { [key: string]: RegExp };

const regex: Regex = {
  required: /./,
  email: /^[\w-\\.]+@([\w-]+\.)+[\w-]{2,4}$/,
  password: /^(?=.*?[A-Z])(?=.*?[a-z])(?=.*?[0-9]).{8,}$/,
  number: / \d+ /
};

export const required: RegExp = /./;

export default regex;
