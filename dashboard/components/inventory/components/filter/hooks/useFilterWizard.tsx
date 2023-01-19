import { NextRouter } from 'next/router';
import { ChangeEvent, useEffect, useState } from 'react';
import { InventoryFilterDataProps } from '../../../hooks/useInventory';

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
  const [data, setData] = useState<InventoryFilterDataProps>(INITIAL_DATA);
  const [costBetween, setCostBetween] =
    useState<CostBetween>(INITIAL_COST_BETWEEN);

  useEffect(() => {
    setData(prev => ({
      ...prev,
      values: [costBetween.min, costBetween.max]
    }));
  }, [costBetween]);

  function resetData() {
    setStep(0);
    setCostBetween(INITIAL_COST_BETWEEN);
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

  function handleCostBetween(newValue: Partial<CostBetween>) {
    setCostBetween(prev => ({ ...prev, ...newValue }));
  }

  function filter() {
    if (router.asPath === '/') {
      router.push(
        `/?${data.field === 'tag' ? `tag:${data.tagKey}` : data.field}:${
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
    data,
    resetData,
    cleanValues,
    filter
  };
}

export default useFilterWizard;
