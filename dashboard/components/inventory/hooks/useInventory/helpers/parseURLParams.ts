import { InventoryFilterData } from '../types/useInventoryTypes';

/** Parse the URL Params.
 * - Argument of type 'fetch' will output the object to fetch an inventory list and stats based on filters.
 * - Argument of type 'display' will output the object to populate the InventoryFilterSummary component
 * Input:
 *  one portion of the URL params: e.g. tag:A:IS_EMPTY
 * */
function parseURLParams(
  param: string | InventoryFilterData,
  type: 'fetch' | 'display',
  view?: boolean
) {
  let formatString;
  let filter;

  if (!view) {
    formatString = (param as string).split(':');
  } else {
    formatString = Object.values(param);
    formatString = [...formatString.slice(0, 2), formatString[2]!.toString()];
  }

  if (formatString[0]!.includes('tag:')) {
    const tag = (formatString[0] as string).split(':');
    formatString = [
      tag[0],
      tag[1],
      formatString[1],
      formatString[2]?.toString()
    ];
  }

  if (formatString[0] === 'tag' && type === 'fetch') {
    if (formatString.length > 2) {
      if (
        formatString.indexOf('IS_EMPTY') !== -1 ||
        formatString.indexOf('IS_NOT_EMPTY') !== -1 ||
        formatString.indexOf('EXISTS') !== -1 ||
        formatString.indexOf('NOT_EXISTS') !== -1
      ) {
        const key = formatString.slice(1, formatString.length - 1).join(':');

        filter = {
          field: `${formatString[0]}:${key}`,
          operator: formatString[formatString.length - 1],
          values: []
        };
      } else {
        const key = formatString.slice(1, formatString.length - 2).join(':');

        filter = {
          field: `${formatString[0]}:${key}`,
          operator: formatString[formatString.length - 2],
          values: (formatString[formatString.length - 1] as string).split(',')
        };
      }
    } else {
      filter = {
        field: `${formatString[0]}:${formatString[1]}`,
        operator: formatString[2],
        values:
          formatString[2] === 'IS_EMPTY' || formatString[2] === 'IS_NOT_EMPTY'
            ? []
            : (formatString[3] as string).split(',')
      };
    }
  }

  if (formatString[0] !== 'tag' && type === 'fetch') {
    filter = {
      field: formatString[0],
      operator: formatString[1],
      values:
        formatString[1] === 'IS_EMPTY' || formatString[1] === 'IS_NOT_EMPTY'
          ? []
          : (formatString[2] as string).split(',')
    };
  }

  if (formatString[0] === 'tag' && type === 'display') {
    if (formatString.length > 2) {
      if (
        formatString.indexOf('IS_EMPTY') !== -1 ||
        formatString.indexOf('IS_NOT_EMPTY') !== -1 ||
        formatString.indexOf('EXISTS') !== -1 ||
        formatString.indexOf('NOT_EXISTS') !== -1
      ) {
        filter = {
          field: formatString[0],
          tagKey: formatString[1],
          operator: formatString[2],
          values: []
        };
      } else {
        const key = formatString.slice(1, formatString.length - 2).join(':');

        filter = {
          field: formatString[0],
          tagKey: key,
          operator: formatString[formatString.length - 2],
          values: (formatString[formatString.length - 1] as string).split(',')
        };
      }
    } else {
      filter = {
        field: formatString[0],
        tagKey: formatString[1],
        operator: formatString[2],
        values:
          formatString[2] === 'IS_EMPTY' || formatString[2] === 'IS_NOT_EMPTY'
            ? []
            : (formatString[3] as string).split(',')
      };
    }
  }

  if (formatString[0] !== 'tag' && type === 'display') {
    filter = {
      field: formatString[0],
      operator: formatString[1],
      values:
        formatString[1] === 'IS_EMPTY' || formatString[1] === 'IS_NOT_EMPTY'
          ? []
          : (formatString[2] as string).split(',')
    };
  }

  return filter as InventoryFilterData;
}

export default parseURLParams;
