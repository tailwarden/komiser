import Button from '../button/Button';

type SidepanelHeaderProps = {
  title: string;
  subtitle: string;
  closeModal: () => void;
  deleteAction?: () => void;
  deleteLabel?: string;
};

function SidepanelHeader({
  title,
  subtitle,
  closeModal,
  deleteAction,
  deleteLabel
}: SidepanelHeaderProps) {
  return (
    <div className="flex flex-wrap-reverse items-center justify-between gap-6 sm:flex-nowrap">
      <div className="flex flex-wrap items-center gap-4 sm:flex-nowrap">
        <div className="flex flex-col gap-1 py-1">
          <p className="... w-48 truncate font-medium text-black-900">
            {title}
          </p>
          <p className="flex items-center gap-2 text-xs text-black-300">
            {subtitle}
          </p>
        </div>
      </div>

      <div className="flex flex-shrink-0 items-center gap-4">
        {deleteAction && (
          <Button style="delete" onClick={deleteAction}>
            {deleteLabel || 'Delete'}
          </Button>
        )}

        <Button style="secondary" onClick={closeModal}>
          Close
        </Button>
      </div>
    </div>
  );
}

export default SidepanelHeader;
