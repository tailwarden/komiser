import { useRouter } from 'next/router';
import { ChangeEvent, useState } from 'react';
import settingsService from '../../../services/settingsService';
import { ToastProps } from '../../toast/hooks/useToast';
import { InventoryItem } from './useInventory';

export type InventoryFilterDataProps = {
  field:
    | 'provider'
    | 'region'
    | 'account'
    | 'name'
    | 'service'
    | string
    | undefined;
  operator:
    | 'IS'
    | 'IS_NOT'
    | 'CONTAINS'
    | 'NOT_CONTAINS'
    | 'IS_EMPTY'
    | 'IS_NOT_EMPTY'
    | undefined;
  tagKey: string | undefined;
  values: [] | string[];
};

const INITIAL_DATA = {
  field: undefined,
  operator: undefined,
  tagKey: '',
  values: []
};

type InventoryFilterProps = {
  applyFilteredInventory: (inventory: InventoryItem[]) => void;
  setToast: (toast: ToastProps) => void;
};

function useFilterWizard({
  applyFilteredInventory,
  setToast
}: InventoryFilterProps) {
  const [step, setStep] = useState(0);
  const [isOpen, setIsOpen] = useState(false);
  const [data, setData] = useState<InventoryFilterDataProps>(INITIAL_DATA);
  const [loading, setLoading] = useState(false);
  const router = useRouter();

  function resetData() {
    setStep(0);
    setData({ ...INITIAL_DATA, values: [] });
  }

  function toggle() {
    resetData();
    setIsOpen(!isOpen);
  }

  function goTo(index: number) {
    if (index === 0) {
      resetData();
    } else {
      setStep(index);
    }
  }

  function cleanValues() {
    setData({ ...data, values: [] });
  }

  function handleField(field: string) {
    setData(prev => ({ ...prev, field }));
    setStep(1);
  }

  function handleOperator(operator: InventoryFilterDataProps['operator']) {
    setData(prev => ({ ...prev, operator }));
    setStep(2);
  }

  function handleTagKey(newValue: { tagKey: string }) {
    setData(prev => ({ ...prev, tagKey: newValue.tagKey }));
  }

  function handleValueCheck(
    e: ChangeEvent<HTMLInputElement>,
    newValue: string
  ) {
    const newValues: string[] = data.values;

    if (e.currentTarget.checked) {
      newValues.push(newValue);
      setData(prev => ({ ...prev, values: newValues }));
    } else {
      const index = newValues.findIndex(value => value === newValue);
      newValues.splice(index, 1);
      setData(prev => ({ ...prev, values: newValues }));
    }
  }

  function handleValueInput(newValue: { values: string }) {
    setData(prev => ({ ...prev, values: [newValue.values] }));
  }

  function filter() {
    setLoading(true);
    const payload = { ...data };

    if (payload.tagKey) {
      payload.field = `tag:${payload.tagKey}`;
    }

    delete payload.tagKey;
    console.log(payload);
    const payloadJson = JSON.stringify([payload]);

    settingsService.getFilteredInventory(payloadJson).then(res => {
      if (res.error) {
        setToast({
          hasError: true,
          title: `Filter could not be applied!`,
          message: `Please refresh the page and try again.`
        });
        setLoading(false);
      } else {
        setToast({
          hasError: false,
          title: `Filter applied!`,
          message: `The filter selection was successfully applied.`
        });
        applyFilteredInventory(res);
        router.push(
          `/?field=${payload.field}&operator=${payload.operator}${
            payload.values.length > 0
              ? `&values=${payload.values.map(value => value)}`
              : ''
          }`,
          undefined,
          { shallow: true }
        );
        setLoading(false);
        toggle();
      }
    });
  }

  return {
    toggle,
    isOpen,
    step,
    goTo,
    handleField,
    handleOperator,
    handleTagKey,
    handleValueCheck,
    handleValueInput,
    data,
    resetData,
    cleanValues,
    filter,
    loading
  };
}

export default useFilterWizard;
