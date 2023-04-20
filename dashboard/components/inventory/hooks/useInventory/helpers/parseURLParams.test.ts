import parseURLParams from './parseURLParams';

describe('parseURLParams', () => {
  it('should return a filter object', () => {
    const param = 'tag:tagKey:tagValue:IS:tagValue';
    const type = 'fetch';
    const filter = parseURLParams(param, type);

    expect(filter).toEqual({
      field: 'tag:tagKey:tagValue',
      operator: 'IS',
      values: ['tagValue']
    });
  });

  it('should return a filter object with a NOT operator', () => {
    const param = 'tag:tagKey:tagValue:NOT:tagValue';
    const type = 'fetch';
    const filter = parseURLParams(param, type);

    expect(filter).toEqual({
      field: 'tag:tagKey:tagValue',
      operator: 'NOT',
      values: ['tagValue']
    });
  });

  it('should return a filter object for the EXISTS tag operator', () => {
    const param = 'tag:tagKey:EXISTS';
    const type = 'display';
    const filter = parseURLParams(param, type);

    expect(filter).toEqual({
      field: 'tag',
      operator: 'EXISTS',
      tagKey: 'tagKey',
      values: []
    });
  });

  it('should return a filter object for the EXISTS tag operator', () => {
    const param = 'tag:tagKey:NOT_EXISTS';
    const type = 'display';
    const filter = parseURLParams(param, type);

    expect(filter).toEqual({
      field: 'tag',
      operator: 'NOT_EXISTS',
      tagKey: 'tagKey',
      values: []
    });
  });
});
