import { motion } from 'framer-motion';
import React from 'react';
import { ArcherElement } from 'react-archer';

interface WordElementProps {
    content: string,
    targetId?: string | string[],
    elementId: string
}

export default function WordElement(WordElementProps: WordElementProps) {
    let relations: any[] = []
    if (WordElementProps.targetId) {
        relations.push({
            targetId: WordElementProps.targetId,
            sourceAnchor: 'right',
            targetAnchor: 'left',
        })
    }

    return (
        <div className='word-container'>
            <motion.div
                layout
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                exit={{ opacity: 0, scale: 0.5 }}
                transition={{ type: "spring", stiffness: 100, damping: 30 }}
                
            >
                <ArcherElement
                    id={WordElementProps.elementId}
                    relations={relations}
                >
                    <p>
                        {WordElementProps.content}
                    </p>
                </ArcherElement>
            </motion.div>
        </div>
    )
}