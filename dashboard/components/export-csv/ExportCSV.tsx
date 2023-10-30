import { ToastProps } from '@components/toast/Toast';
import { useRouter } from 'next/router';
import settingsService from '../../services/settingsService';
import ExportCSVButton from './ExportCSVButton';

type ExportCSVProps = {
  showToast: (toast: ToastProps) => void;
};

function ExportCSV({ showToast }: ExportCSVProps) {
  const router = useRouter();

  function exportCSV(id?: string) {
    settingsService.exportCSV(id);
    showToast({
      hasError: false,
      title: 'CSV exported',
      message: 'The download of the CSV file should begin shortly.'
    });
  }

  const isFilteredList =
    Object.keys(router.query).length > 0 && !router.query.view;
  const id = router.query.view ? router.query.view.toString() : undefined;

  return (
    <ExportCSVButton id={id} disabled={isFilteredList} exportCSV={exportCSV} />
  );
}

export default ExportCSV;
