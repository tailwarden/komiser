import { ChangeEvent, useEffect, useState } from 'react';
import settingsService from '../../../../services/settingsService';
import regex from '../../../../utils/regex';
import Checkbox from '../../../checkbox/Checkbox';
import Input from '../../../input/Input';
import { InventoryFilterDataProps } from '../../hooks/useFilterWizard';

type InventoryFilterValueProps = {
  data: InventoryFilterDataProps;
  handleValueCheck: (
    e: ChangeEvent<HTMLInputElement>,
    newValue: string
  ) => void;
  handleValueInput: (newValue: { values: string }) => void;
  cleanValues: () => void;
};

type Options = string[];

function InventoryFilterValue({
  data,
  handleValueCheck,
  handleValueInput,
  cleanValues
}: InventoryFilterValueProps) {
  const [options, setOptions] = useState<Options | undefined>();
  const [error, setError] = useState(false);

  useEffect(() => {
    let mounted = true;

    if (data.operator === 'IS_EMPTY' || data.operator === 'IS_NOT_EMPTY') {
      cleanValues();
      setOptions(undefined);
    } else {
      if (data.field === 'provider') {
        settingsService.getProviders().then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setOptions(res);
            }
          }
        });
      }

      if (data.field === 'account') {
        settingsService.getAccounts().then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setOptions(res);
            }
          }
        });
      }

      if (data.field === 'region') {
        settingsService.getRegions().then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setOptions(res);
            }
          }
        });
      }

      if (data.field === 'service') {
        settingsService.getServices().then(res => {
          if (mounted) {
            if (res === Error) {
              setError(true);
            } else {
              setOptions(res);
            }
          }
        });
      }
    }

    return () => {
      mounted = false;
    };
  }, []);

  return (
    <div className="flex flex-col gap-2 min-w-[20rem]">
      {/* Display multi-select */}
      {options &&
        options.length > 0 &&
        options.map((option, idx) => (
          <div key={idx} className="flex items-center gap-2 py-1">
            <Checkbox
              id={option}
              onChange={e => handleValueCheck(e, option)}
              checked={!!data.values.find(value => value === option)}
            />
            <label htmlFor={option} className="w-full">
              {option}
            </label>
          </div>
        ))}

      {/* Display input for resource name and tag values */}
      {!options &&
        data.operator !== 'IS_EMPTY' &&
        data.operator !== 'IS_NOT_EMPTY' && (
          <div className="pl-1 pt-2 pr-4 pb-2">
            <Input
              type="text"
              name="values"
              label={data.field === 'tag' ? 'Tag value' : 'Resource name'}
              value={data.values}
              regex={regex.required}
              error="Please provide a value"
              action={handleValueInput}
            />
          </div>
        )}
    </div>
  );
}

export default InventoryFilterValue;
