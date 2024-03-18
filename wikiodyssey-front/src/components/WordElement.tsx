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
            <ArcherElement
                id={WordElementProps.elementId}
                relations={relations}
            >
                <p>
                    {WordElementProps.content}
                </p>
            </ArcherElement>
        </div>
    )
}