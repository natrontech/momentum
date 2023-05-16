import React from "react";

interface ButtonProps {
    text: string;
    onClick: () => void;
}

const Button: React.FC<ButtonProps> = ({ text, onClick }) => {
    return (
        <button
            onClick={onClick}
            className="px-6 py-3 bg-blue-600 text-white font-semibold rounded shadow-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
            {text}
        </button>
    );
};

export default Button;
