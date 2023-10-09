import { useState, useRef, useCallback, memo, SyntheticEvent } from 'react';
import { FileUploader } from 'react-drag-drop-files';
// eslint-disable-next-line import/no-extraneous-dependencies
import { toBlob } from 'html-to-image';
import Image from 'next/image';

import Modal from '@components/modal/Modal';
import Input from '@components/input/Input';
import settingsService from '@services/settingsService';
import Button from '@components/button/Button';
import useToast from '@components/toast/hooks/useToast';
import Toast from '@components/toast/Toast';

// We define the placeholder here for convenience
// It's difficult to read when passed inline
const textAreaPlaceholder = `Example:
Steps to Reproduce
1. Describe the actions you took leading to the bug.
2. Include any specific settings or options you selected.

Expected Behavior
1. Explain what you expected to happen.
2. Detail how the feature or function should work.

Outcome
1. Describe what actually happened.
2. Include any error messages or unexpected behavior.`;

const useFeedbackWidget = (defaultState: boolean = false) => {
  const [showFeedbackModel, setShowFeedbackModal] = useState(defaultState);

  const FILE_TYPES = ['JPG', 'PNG', 'GIF', 'TXT', 'LOG'];
  const FEEDBACK_MODAL_ID = 'feedback-modal';

  function openFeedbackModal() {
    setShowFeedbackModal(true);
  }

  function closeFeedbackModal() {
    setShowFeedbackModal(false);
  }

  function toggleFeedbackModal() {
    setShowFeedbackModal(!showFeedbackModel);
  }

  const filter = (node: HTMLElement) => !node.id?.startsWith(FEEDBACK_MODAL_ID);

  const FeedbackModal = () => {
    const [email, updateEmail] = useState('');
    const [description, updateDescription] = useState('');
    const [screenshotBlob, setScreenshotBlob] = useState<Blob | null>(null);
    const [fileAttachement, setFileAttachement] = useState<File | null>(null);
    const [dragActive, setDragActive] = useState(false);
    const [isSendingFeedback, setIsSendingFeedback] = useState(false);
    const { toast, setToast, dismissToast } = useToast();

    const takeScreenshot = () => {
      if (document.documentElement === null || isSendingFeedback) {
        return;
      }

      toBlob(document.documentElement, { cacheBust: true, filter })
        .then(blob => {
          setScreenshotBlob(blob);
          setToast({
            hasError: false,
            title: 'Screen capture',
            message:
              'A screenshot of your current page on Komiser has been captured and attached to your feedback.'
          });
        })
        .catch(err => {
          setToast({
            hasError: true,
            title: 'Screen capture failed',
            message:
              'The capture of your current page on Komiser couldn’t be saved. Please try again or upload a screenshot manually. Our support is also happy to help you!'
          });
        });
    };

    function clearFeedbackForm() {
      setScreenshotBlob(null);
      setFileAttachement(null);
      updateDescription('');
      updateEmail('');
    }

    async function uploadFeedback(e: SyntheticEvent) {
      if (!isSendingFeedback) {
        try {
          setIsSendingFeedback(true);
          e.preventDefault();
          const formData = new FormData();

          formData.append('description', description);
          if (email) formData.append('email', email);
          if (screenshotBlob && screenshotBlob !== null)
            formData.append('image', screenshotBlob);
          if (
            screenshotBlob === null &&
            fileAttachement &&
            fileAttachement !== null
          )
            formData.append('image', fileAttachement);

          settingsService
            .sendFeedback(formData)
            .then(result => {
              setToast({
                hasError: false,
                title: 'Feedback sent',
                message:
                  result.Response ||
                  'Your insights are valuable to us. Thank you for helping us improve!'
              });
              clearFeedbackForm();
            })
            .catch(error => {
              setToast({
                hasError: true,
                title: 'Feedback',
                message: 'An Error happened. Maybe try again please!'
              });
            })
            .finally(() => {
              setIsSendingFeedback(false);
            });
        } catch {
          setIsSendingFeedback(false);
        }
      }
    }

    const uploadFile = (attachement: File) => {
      setFileAttachement(attachement);
    };

    return (
      <>
        <Modal
          isOpen={showFeedbackModel}
          closeModal={() => closeFeedbackModal()}
          id={FEEDBACK_MODAL_ID}
        >
          <div className="w-96">
            <h3 className="text-lg font-bold text-black-900">
              Describe your issue
            </h3>
            <p className="text-base text-black-400">
              By providing details of the issue you’ve encountered and outlining
              the steps to reproduce it, we’ll be able to give you better
              support.
            </p>
            <form onSubmit={uploadFeedback} className="mt-4">
              <Input
                disabled={isSendingFeedback}
                type="email"
                label="Your email (optional)"
                name="email"
                action={change => {
                  updateEmail(change.email);
                }}
                value={email}
              />
              <textarea
                disabled={isSendingFeedback}
                rows={15}
                placeholder={textAreaPlaceholder}
                className="peer mt-4 w-full rounded bg-white px-4 pb-[0.75rem] pt-[1.75rem] text-sm text-black-900 caret-primary outline outline-[0.063rem] outline-black-200 focus:outline-[0.12rem] focus:outline-primary"
                onChange={event => updateDescription(event?.target?.value)}
                value={description}
                required
              />
              <div className="mt-4 flex justify-between gap-2">
                {(!screenshotBlob || screenshotBlob === null) && (
                  <a
                    className="w-[50%] cursor-pointer border border-black-170 px-6 py-5 text-center text-xs transition hover:border-primary"
                    onClick={() => takeScreenshot()}
                  >
                    <svg
                      className="m-4 inline-block"
                      width="25"
                      height="24"
                      viewBox="0 0 25 24"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M2.5 9V6.5C2.5 4.01 4.51 2 7 2H9.5"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                      <path
                        d="M15.5 2H18C20.49 2 22.5 4.01 22.5 6.5V9"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                      <path
                        d="M22.5 16V17.5C22.5 19.99 20.49 22 18 22H16.5"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                      <path
                        d="M9.5 22H7C4.51 22 2.5 19.99 2.5 17.5V15"
                        stroke="#0C1717"
                        stroke-width="1.5"
                        stroke-linecap="round"
                        stroke-linejoin="round"
                      />
                    </svg>
                    <br />
                    Capture current screen
                  </a>
                )}
                {screenshotBlob && screenshotBlob !== null && (
                  <div className="relative h-[76px] w-[217px]">
                    <Image
                      src={URL.createObjectURL(screenshotBlob)}
                      alt="attached screenshot"
                      width="217"
                      height="76"
                    />
                    <a
                      onClick={() =>
                        !isSendingFeedback && setScreenshotBlob(null)
                      }
                      className="absolute right-4 top-4 cursor-pointer"
                    >
                      x
                    </a>
                  </div>
                )}
                <div className="relative">
                  <FileUploader
                    handleChange={uploadFile}
                    name="attachement"
                    types={FILE_TYPES}
                    fileOrFiles={fileAttachement}
                    onTypeError={(err: string) =>
                      setToast({
                        hasError: true,
                        title: 'File upload failed',
                        message: err
                      })
                    }
                    onSizeError={(err: string) =>
                      setToast({
                        hasError: true,
                        title: 'File upload failed',
                        message: err
                      })
                    }
                  >
                    {fileAttachement === null && (
                      <div className="cursor-pointer border border-dashed border-black-170 px-6 py-5 text-center text-xs transition hover:border-primary">
                        <svg
                          className="m-4 inline-block"
                          width="25"
                          height="24"
                          viewBox="0 0 25 24"
                          fill="none"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                            d="M9.5 17V11L7.5 13"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M9.5 11L11.5 13"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M22.5 10V15C22.5 20 20.5 22 15.5 22H9.5C4.5 22 2.5 20 2.5 15V9C2.5 4 4.5 2 9.5 2H14.5"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M22.5 10H18.5C15.5 10 14.5 9 14.5 6V2L22.5 10Z"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                        </svg>
                        <br />
                        Drag and drop or chose a file
                      </div>
                    )}
                    {fileAttachement !== null && (
                      <div className="relative h-[76px] w-[217px]">
                        <svg
                          width="40"
                          height="40"
                          viewBox="0 0 40 40"
                          fill="none"
                          xmlns="http://www.w3.org/2000/svg"
                        >
                          <path
                            d="M35 11.6666V28.3333C35 33.3333 32.5 36.6666 26.6667 36.6666H13.3333C7.5 36.6666 5 33.3333 5 28.3333V11.6666C5 6.66659 7.5 3.33325 13.3333 3.33325H26.6667C32.5 3.33325 35 6.66659 35 11.6666Z"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-miterlimit="10"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M24.167 7.5V10.8333C24.167 12.6667 25.667 14.1667 27.5003 14.1667H30.8337"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-miterlimit="10"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M13.333 21.6667H19.9997"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-miterlimit="10"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                          <path
                            d="M13.333 28.3333H26.6663"
                            stroke="#0C1717"
                            stroke-width="1.5"
                            stroke-miterlimit="10"
                            stroke-linecap="round"
                            stroke-linejoin="round"
                          />
                        </svg>
                        <a
                          onClick={e => {
                            e.preventDefault();
                            if (!isSendingFeedback) setFileAttachement(null);
                            return false;
                          }}
                          className="absolute right-4 top-4 cursor-pointer"
                        >
                          x
                        </a>
                      </div>
                    )}
                  </FileUploader>
                </div>
              </div>
              <div className="mt-4 flex justify-between">
                <p className="text-xs text-black-400">
                  Need in depth assistance?
                  <br />
                  Email us at{' '}
                  <a href="mailto:support@tailwarden.com">
                    support@tailwarden.com
                  </a>
                  .
                </p>
                <Button type="submit" disabled={isSendingFeedback}>
                  {isSendingFeedback ? `Sending...` : `Send Feedback`}
                </Button>
              </div>
            </form>
          </div>
        </Modal>
        {toast && <Toast {...toast} dismissToast={dismissToast} />}
      </>
    );
  };

  return {
    openFeedbackModal,
    closeFeedbackModal,
    toggleFeedbackModal,
    FeedbackModal: memo(FeedbackModal)
  };
};

export default useFeedbackWidget;
