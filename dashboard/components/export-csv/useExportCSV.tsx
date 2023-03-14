import { useState } from 'react';
import settingsService from '../../services/settingsService';
import { ToastProps } from '../toast/hooks/useToast';

type useExportCSVProps = {
  setToast: (toast: ToastProps | undefined) => void;
};

function useExportCSV({ setToast }: useExportCSVProps) {
  const [loading, setLoading] = useState(false);

  function exportCSV(id?: string) {
    setLoading(true);

    settingsService.exportCSV(id || undefined).then(res => {
      if (res === Error) {
        setLoading(false);
        setToast({
          hasError: true,
          title: 'CSV not exported',
          message:
            'There was an error exporting the CSV for this list of resources.'
        });
      } else {
        setToast({
          hasError: false,
          title: 'CSV exported',
          message: 'The download of the CSV file should begin shortly.'
        });
        setLoading(false);
      }
    });
  }

  return { loading, exportCSV };
}

export default useExportCSV;
