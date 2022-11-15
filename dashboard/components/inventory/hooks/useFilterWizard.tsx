import { useState } from 'react';

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
    | 'EMPTY'
    | 'NOT_EMPTY'
    | undefined;
  values: null | string[];
};

const INITIAL_DATA = {
  field: undefined,
  operator: undefined,
  values: null
};

function useFilterWizard() {
  const [step, setStep] = useState(0);
  const [isOpen, setIsOpen] = useState(false);
  const [data, setData] = useState<InventoryFilterDataProps>(INITIAL_DATA);

  console.log(data);

  function resetData() {
    setStep(0);
    setData(INITIAL_DATA);
  }

  function toggle() {
    resetData();
    setIsOpen(!isOpen);
  }

  function goTo(index: number) {
    setStep(index);
  }

  function handleField(field: string) {
    setData(prev => ({ ...prev, field }));
    setStep(1);
  }

  function handleOperator(operator: InventoryFilterDataProps['operator']) {
    setData(prev => ({ ...prev, operator }));
    setStep(2);
  }

  return {
    toggle,
    isOpen,
    step,
    goTo,
    handleField,
    handleOperator,
    data,
    resetData
  };
}

export default useFilterWizard;
