import { ToastProps } from '../toast/hooks/useToast';
import ExportCSVButton from './ExportCSVButton';
import useExportCSV from './useExportCSV';

type ExportCSVProps = {
  id?: string;
  setToast: (toast: ToastProps | undefined) => void;
};

function ExportCSV({ id, setToast }: ExportCSVProps) {
  const { loading, exportCSV } = useExportCSV({ setToast });

  return <ExportCSVButton id={id} loading={loading} exportCSV={exportCSV} />;
}

export default ExportCSV;
