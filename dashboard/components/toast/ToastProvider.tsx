import React, {
  createContext,
  useContext,
  useState,
  useEffect,
  FC,
  ReactNode
} from 'react';

export type ToastProps = {
  hasError: boolean;
  title: string;
  message: string;
};

type ToastContextType = {
  showToast: (newToast: ToastProps) => void;
  dismissToast: () => void;
  toast: ToastProps | null;
};

const ToastContext = createContext<ToastContextType | undefined>(undefined);

export const ToastProvider: FC<{ children: ReactNode }> = ({ children }) => {
  const [toast, setToast] = useState<ToastProps | null>(null);

  const dismissToast = () => {
    setToast(null);
  };
  const showToast = (newToast: ToastProps) => {
    setToast(newToast);
  };

  useEffect(() => {
    let timeout: any;
    if (toast) {
      timeout = setTimeout(dismissToast, 5000);
    }
    return () => clearTimeout(timeout);
  }, [toast]);

  return (
    <ToastContext.Provider value={{ toast, showToast, dismissToast }}>
      {children}
    </ToastContext.Provider>
  );
};

export const useToast = () => {
  const context = useContext(ToastContext);

  if (context === undefined) {
    throw new Error('useToast must be used within a ToastProvider');
  }

  return context;
};
