import { InventoryFilterData } from '../types/useInventoryTypes';
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

  it('should return a filter object for the EXISTS tag operator if parsing for InventoryFilterData (view = true)', () => {
    const param: InventoryFilterData = {
      field: 'tag:tagKey',
      operator: 'EXISTS',
      values: []
    };
    const type = 'display';
    const filter = parseURLParams(param, type, true);

    expect(filter).toEqual({
      field: 'tag',
      operator: 'EXISTS',
      tagKey: 'tagKey',
      values: []
    });
  });

  it('should return a filter object for the IS_NOT tag operator if parsing for InventoryFilterData (view = true)', () => {
    const param: InventoryFilterData = {
      field: 'tag:tagKey',
      operator: 'IS_NOT',
      values: ['tagValue']
    };
    const type = 'display';
    const filter = parseURLParams(param, type, true);

    expect(filter).toEqual({
      field: 'tag',
      operator: 'IS_NOT',
      tagKey: 'tagKey',
      values: ['tagValue']
    });
  });

  it('should return a filter object for non-tag operators', () => {
    const param: InventoryFilterData = {
      field: 'Cost',
      operator: 'IS_GREATER_THAN',
      values: ['100']
    };
    const type = 'display';
    const filter = parseURLParams(param, type, true);

    expect(filter).toEqual({
      field: 'Cost',
      operator: 'IS_GREATER_THAN',
      values: ['100']
    });
  });

  it('should return a filter object for non-tag operators #2', () => {
    const param: InventoryFilterData = {
      field: 'Cost',
      operator: 'IS_BETWEEN',
      values: ['100', '200']
    };
    const type = 'display';
    const filter = parseURLParams(param, type, true);

    expect(filter).toEqual({
      field: 'Cost',
      operator: 'IS_BETWEEN',
      values: ['100', '200']
    });
  });
});
