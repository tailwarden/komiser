import { ToastProps } from '../toast/hooks/useToast';
import ExportCSVButton from './ExportCSVButton';
import useExportCSV from './useExportCSV';

type ExportCSVProps = {
  displayInTable?: boolean;
  setToast: (toast: ToastProps | undefined) => void;
};

function ExportCSV({ displayInTable = false, setToast }: ExportCSVProps) {
  const { id, isFilteredList, loading, exportCSV } = useExportCSV({ setToast });

  return (
    <ExportCSVButton
      id={id}
      loading={loading}
      disabled={isFilteredList}
      displayInTable={displayInTable}
      exportCSV={exportCSV}
    />
  );
}

export default ExportCSV;
