import Button from '../button/Button';
import DownloadIcon from '../icons/DownloadIcon';

type ExportCSVButtonProps = {
  id?: string;
  loading: boolean;
  exportCSV: (id?: string) => void;
};

function ExportCSVButton({ id, loading, exportCSV }: ExportCSVButtonProps) {
  return (
    <Button
      style="ghost"
      size="sm"
      align="left"
      gap="md"
      transition={false}
      loading={loading}
      onClick={() => exportCSV(id || undefined)}
    >
      {!loading && <DownloadIcon width={24} height={24} />}
      Download CSV
    </Button>
  );
}

export default ExportCSVButton;
