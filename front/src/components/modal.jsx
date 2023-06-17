import { Transition } from '@headlessui/react';
import { Children, Fragment } from 'react';

const Modal = ({ children, icon, title, isOpen, onClose, className }) => {

    let _body;

    Children.forEach(children, child => {
        if (child.type === ModalBody) {
            return _body = child
        }
    })

    return (
        <Transition 
            show={isOpen}
            enter="transition duration-500 ease-out"
            enterFrom="transform opacity-0"
            enterTo="transform opacity-500"
            leave="transition duration-300 ease-out"
            leaveFrom="transform opacity-100"
            leaveTo="transform opacity-0"
            as={Fragment}
        >
            <div className={`absolute top-[50%] left-[50%] translate-x-[-50%] translate-y-[-50%] max-w-[386px] w-full max-h-[240px] h-full z-50 ${className}`}>
                <div className="flex flex-col w-full h-full drop-shadow-modal modal-body-bg">
                    {/* Header */}
                    <div className="flex justify-between w-full h-[27px]">
                        <div className="flex items-center w-auto ml-[9px]">
                            <div className={`bg-center bg-no-repeat bg-contain ${icon} w-[14px] h-[14px]`} />
                            <div className="ml-[4px] text-[12px] modal-text-shadow leading-[16px]">{title}</div>
                        </div>
                        <div className="w-auto z-0">
                            <div className="relative bg-center bg-no-repeat bg-contain bg-close-inactive hover:bg-close-active transition-all ease-in-out cursor-pointer w-[43px] h-[18px] modal-close-button mt-[2px] mr-[6px]" onClick={onClose} />
                        </div>
                    </div>
                    {/* Body */}
                    {_body}
                </div>
            </div>
        </Transition>
    )
}

{/* Body */}
const ModalBody = ({ children }) => {
    return (
        <div className="h-full w-full px-[7px] pb-[7px] pt-[4px]">
            <div className="flex items-center bg-white w-full h-full">
                {children}
            </div>
        </div>
    )
}

Modal.Body = ModalBody


export default Modal