import { ReactNode, useState } from 'react';

function useFilterWizard(steps: ReactNode[]) {
  const [currentStep, setCurrentStep] = useState(0);

  function back() {
    setCurrentStep(prev => {
      if (prev === 0) {
        return prev;
      }
      return prev - 1;
    });
  }

  function next() {
    setCurrentStep(prev => {
      if (prev >= steps.length - 1) {
        return prev;
      }
      return prev + 1;
    });
  }

  function goTo(index: number) {
    setCurrentStep(index);
  }

  return {
    step: steps[currentStep],
    steps,
    currentStep: currentStep + 1,
    isFirstStep: currentStep === 0,
    isFinishStep: currentStep === steps.length - 2,
    isLastStap: currentStep === steps.length - 1,
    back,
    next,
    goTo
  };
}

export default useFilterWizard;
