import { NextRouter } from 'next/router';
import { ChangeEvent, useEffect, useState } from 'react';
import { InventoryFilterData } from '../../../hooks/useInventory/types/useInventoryTypes';

const INITIAL_DATA = {
  field: undefined,
  operator: undefined,
  tagKey: '',
  values: []
};

const INITIAL_COST_BETWEEN = {
  min: '',
  max: ''
};

const INITIAL_INLINE_ERROR = { hasError: false, message: '' };

type InventoryFilterProps = {
  router: NextRouter;
  setSkippedSearch: (number: number) => void;
};

export type CostBetween = {
  min: string;
  max: string;
};

function useFilterWizard({ router, setSkippedSearch }: InventoryFilterProps) {
  const [step, setStep] = useState(0);
  const [isOpen, setIsOpen] = useState(false);
  const [data, setData] = useState<InventoryFilterData>(INITIAL_DATA);
  const [costBetween, setCostBetween] =
    useState<CostBetween>(INITIAL_COST_BETWEEN);
  const [inlineError, setInlineError] = useState(INITIAL_INLINE_ERROR);

  useEffect(() => {
    if (costBetween.min || costBetween.max) {
      setData(prev => ({
        ...prev,
        values: [costBetween.min, costBetween.max]
      }));
    }
  }, [costBetween]);

  function resetData() {
    setStep(0);
    setCostBetween(INITIAL_COST_BETWEEN);
    setInlineError(INITIAL_INLINE_ERROR);
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

  function handleOperator(operator: InventoryFilterData['operator']) {
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

  function handleCostBetween(newValue: Partial<CostBetween>) {
    setCostBetween(prev => ({ ...prev, ...newValue }));
  }

  function filter() {
    setInlineError(INITIAL_INLINE_ERROR);

    if (data.operator === 'BETWEEN') {
      if (!data.values[0] || !data.values[1]) {
        setInlineError({
          hasError: true,
          message: 'Please provide a min and max value.'
        });
        return null;
      }

      if (Number(data.values[0]) > Number(data.values[1])) {
        setInlineError({
          hasError: true,
          message: 'Max number needs to be higher than min number.'
        });
        return null;
      }

      if (Number(data.values[0]) === Number(data.values[1])) {
        setInlineError({
          hasError: true,
          message: 'Min and max values can not be the same.'
        });
        return null;
      }
    }

    if (router.asPath === '/inventory/') {
      router.push(
        `?${data.field === 'tag' ? `tag:${data.tagKey}` : data.field}:${
          data.operator
        }${
          data.values.length > 0 ? `:${data.values.map(value => value)}` : ''
        }`,
        undefined,
        { shallow: true }
      );
    } else if (router.asPath === '/explorer/') {
      router.push(
        `?${data.field === 'tag' ? `tag:${data.tagKey}` : data.field}:${
          data.operator
        }${
          data.values.length > 0 ? `:${data.values.map(value => value)}` : ''
        }`,
        undefined,
        { shallow: true }
      );
    } else {
      router.push(
        `${router.asPath}&${
          data.field === 'tag' ? `tag:${data.tagKey}` : data.field
        }${`:${data.operator}`}${
          data.values.length > 0 ? `:${data.values.map(value => value)}` : ''
        }`
      );
    }
    setSkippedSearch(0);
    toggle();

    return null;
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
    costBetween,
    handleCostBetween,
    inlineError,
    data,
    resetData,
    cleanValues,
    filter
  };
}

export default useFilterWizard;
