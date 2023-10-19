import { ToastProps } from '@components/toast/Toast';
import { ChangeEvent, useEffect, useState } from 'react';
import settingsService from '../../../../services/settingsService';
import Checkbox from '../../../checkbox/Checkbox';
import Input from '../../../input/Input';
import { InventoryFilterData } from '../../hooks/useInventory/types/useInventoryTypes';
import { CostBetween } from './hooks/useFilterWizard';

type InventoryFilterValueProps = {
  data: InventoryFilterData;
  handleValueCheck: (
    e: ChangeEvent<HTMLInputElement>,
    newValue: string
  ) => void;
  handleValueInput: (newValue: { values: string }) => void;
  cleanValues: () => void;
  showToast: (toast: ToastProps) => void;
  costBetween: CostBetween;
  handleCostBetween: (newValue: Partial<CostBetween>) => void;
};

type Options = string[];

function InventoryFilterValue({
  data,
  handleValueCheck,
  handleValueInput,
  cleanValues,
  showToast,
  costBetween,
  handleCostBetween
}: InventoryFilterValueProps) {
  const [options, setOptions] = useState<Options | undefined>();

  useEffect(() => {
    if (
      data.operator === 'IS_EMPTY' ||
      data.operator === 'IS_NOT_EMPTY' ||
      data.operator === 'EXISTS' ||
      data.operator === 'NOT_EXISTS'
    ) {
      cleanValues();
      setOptions(undefined);
    } else {
      if (data.field === 'provider') {
        settingsService.getProviders().then(res => {
          if (res === Error) {
            showToast({
              hasError: true,
              title: `There was an error when fetching the cloud providers`,
              message: `Please refresh the page and try again.`
            });
          } else {
            setOptions(res);
          }
        });
      }

      if (data.field === 'account') {
        settingsService.getAccounts().then(res => {
          if (res === Error) {
            showToast({
              hasError: true,
              title: `There was an error when fetching the cloud accounts`,
              message: `Please refresh the page and try again.`
            });
          } else {
            setOptions(res);
          }
        });
      }

      if (data.field === 'region') {
        settingsService.getRegions().then(res => {
          if (res === Error) {
            showToast({
              hasError: true,
              title: `There was an error when fetching the cloud regions`,
              message: `Please refresh the page and try again.`
            });
          } else {
            setOptions(res);
          }
        });
      }

      if (data.field === 'service') {
        settingsService.getServices().then(res => {
          if (res === Error) {
            showToast({
              hasError: true,
              title: `There was an error when fetching the cloud services`,
              message: `Please refresh the page and try again.`
            });
          } else {
            setOptions(res);
          }
        });
      }
    }
  }, []);

  return (
    <div className="flex min-w-[19.05rem] flex-col gap-2">
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
        data.field !== 'cost' &&
        data.field !== 'relations' &&
        data.operator !== 'IS_EMPTY' &&
        data.operator !== 'IS_NOT_EMPTY' &&
        data.operator !== 'EXISTS' &&
        data.operator !== 'NOT_EXISTS' && (
          <div className="pb-2 pl-1 pr-4 pt-2">
            <Input
              type="text"
              name="values"
              label={data.field === 'tag' ? 'Tag value' : 'Resource name'}
              value={data.values}
              error="Please provide a value"
              action={handleValueInput}
              autofocus={true}
              maxLength={64}
            />
          </div>
        )}

      {/* Display input for cost when is equal, greater or less than */}
      {!options &&
        (data.field === 'cost' || data.field === 'relations') &&
        data.operator !== 'BETWEEN' && (
          <div className="pb-2 pl-1 pr-4 pt-2">
            <Input
              type="number"
              name="values"
              label="Value"
              value={data.values}
              error="Value must be higher than 0."
              action={handleValueInput}
              autofocus={true}
              min={0}
              positiveNumberOnly={true}
            />
          </div>
        )}

      {/* Display input for cost when is between */}
      {!options && data.field === 'cost' && data.operator === 'BETWEEN' && (
        <div className="pb-2 pl-1 pr-4 pt-2">
          <div className="grid grid-cols-2 gap-4">
            <Input
              type="number"
              name="min"
              label="Min value"
              value={costBetween.min}
              error="Value must be higher than 0."
              action={handleCostBetween}
              autofocus={true}
              min={0}
              positiveNumberOnly={true}
            />
            <Input
              type="number"
              name="max"
              label="Max value"
              value={costBetween.max}
              error="Value must be higher than 0."
              action={handleCostBetween}
              min={0}
              positiveNumberOnly={true}
            />
          </div>
        </div>
      )}
    </div>
  );
}

export default InventoryFilterValue;
