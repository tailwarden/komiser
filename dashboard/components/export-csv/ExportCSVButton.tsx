import Button from '../button/Button';
import DownloadIcon from '../icons/DownloadIcon';
import Tooltip from '../tooltip/Tooltip';

type ExportCSVButtonProps = {
  id?: string;
  loading: boolean;
  disabled: boolean;
  displayInTable: boolean;
  exportCSV: (id?: string) => void;
};

function ExportCSVButton({
  id,
  loading,
  disabled,
  displayInTable,
  exportCSV
}: ExportCSVButtonProps) {
  return (
    <>
      <div className="peer flex flex-col">
        <Button
          style={displayInTable ? 'secondary' : 'ghost'}
          size="sm"
          align="left"
          gap="md"
          transition={false}
          loading={loading}
          disabled={disabled}
          onClick={() => exportCSV(id || undefined)}
        >
          {!loading && <DownloadIcon width={24} height={24} />}
          Download CSV
        </Button>
      </div>
      {disabled && (
        <Tooltip top="sm" align="right" width="lg">
          This feature isn&apos;t available yet. To download data from a
          filtered table, save it as a view and download it from there.
        </Tooltip>
      )}
    </>
  );
}

export default ExportCSVButton;
