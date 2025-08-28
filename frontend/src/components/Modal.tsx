import React from 'react';

type Props = { onClose: () => void };

const Modal: React.FC<Props> = ({ onClose }) => (
  <div className="fixed inset-0 flex items-center justify-center bg-black/50">
    <div className="bg-white p-4">
      <p className="text-black">Lazy Modal</p>
      <button onClick={onClose} className="mt-2">Close</button>
    </div>
  </div>
);

export default Modal;
