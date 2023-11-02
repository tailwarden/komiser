import Button from '@components/button/Button';

export type SidepanelFooterProps = {
  loading?: boolean;
  closeModal: () => void;
  saveAction: () => void;
  saveLabel?: string;
};

function SidepanelFooter({
  closeModal,
  saveAction,
  saveLabel,
  loading
}: SidepanelFooterProps) {
  return (
    <>
      <div className="inline-flex w-full items-center justify-end rounded-bl-lg rounded-br-lg border-t border-gray-200 bg-white px-0 py-3">
        <div className="flex shrink grow basis-0 items-center justify-end gap-6">
          <Button size="md" style="ghost" onClick={closeModal}>
            Cancel
          </Button>
          <Button
            size="md"
            type="submit"
            loading={loading}
            onClick={saveAction}
          >
            {saveLabel || 'Save'}
          </Button>
        </div>
      </div>
    </>
  );
}

export default SidepanelFooter;
