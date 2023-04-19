import { useRouter } from 'next/router';
import settingsService from '../../services/settingsService';
import { ToastProps } from '../toast/hooks/useToast';
import ExportCSVButton from './ExportCSVButton';

type ExportCSVProps = {
  setToast: (toast: ToastProps | undefined) => void;
};

function ExportCSV({ setToast }: ExportCSVProps) {
  const router = useRouter();

  function exportCSV(id?: string) {
    settingsService.exportCSV(id);
    setToast({
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
