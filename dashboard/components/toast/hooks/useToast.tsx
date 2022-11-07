import { useEffect, useState } from 'react';

export type ToastProps = {
  hasError: boolean;
  title: string;
  message: string;
};

function useToast() {
  const [toast, setToast] = useState<ToastProps | undefined>(undefined);

  function dismissToast() {
    setToast(undefined);
  }

  useEffect(() => {
    const timeout = setTimeout(dismissToast, 5000);
    return () => {
      clearTimeout(timeout);
    };
  }, [toast]);

  return {
    toast,
    setToast,
    dismissToast
  };
}

export default useToast;
