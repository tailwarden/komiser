import Button from '../button/Button';

type SidepanelHeaderProps = {
  title: string;
  subtitle: string;
  closeModal: () => void;
  deleteAction?: () => void;
};

function SidepanelHeader({
  title,
  subtitle,
  closeModal,
  deleteAction
}: SidepanelHeaderProps) {
  return (
    <div className="flex flex-wrap-reverse sm:flex-nowrap items-center justify-between gap-6">
      <div className="flex flex-wrap sm:flex-nowrap items-center gap-4">
        <div className="flex flex-col gap-1 py-1">
          <p className="font-medium text-black-900 w-48 truncate ...">
            {title}
          </p>
          <p className="flex items-center gap-2 text-xs text-black-300">
            {subtitle}
          </p>
        </div>
      </div>

      <div className="flex items-center gap-4 flex-shrink-0">
        {deleteAction && (
          <Button style="delete-ghost" onClick={deleteAction}>
            Delete
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
