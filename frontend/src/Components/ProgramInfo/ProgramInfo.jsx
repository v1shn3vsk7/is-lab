import React, { useState } from "react";
import './ProgramInfo.css'

const ProgramInfo = () => {
    const [isOpen, setIsOpen] = useState(false);

    const handleClick = () => {
        setIsOpen(!isOpen);
    };

    const ModalContent = () => (
        <div className="modal-content">
            <h2>О Программе</h2>
            <p>Автор: Васильев Владимир</p>
            <p>Вариант: №5</p>
        </div>
    );

    return (
        <div className="program-info">
            <h3 onClick={handleClick}>Справка</h3>
            {isOpen && <ModalContent />}
        </div>
    );
};

export default ProgramInfo;

