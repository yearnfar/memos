import clsx from "clsx";
import React from "react";
import { Resource } from "@/types/proto/api/v1/resource_service";
import { getResourceType, getResourceUrl } from "@/utils/resource";
import Icon from "./Icon";
import showPreviewImageDialog from "./PreviewImageDialog";
import SquareDiv from "./kit/SquareDiv";

interface Props {
  resource: Resource;
  className?: string;
  strokeWidth?: number;
}

const ResourceIcon = (props: Props) => {
  const { resource } = props;
  const resourceType = getResourceType(resource);
  const resourceUrl = getResourceUrl(resource);
  const className = clsx("w-full h-auto", props.className);
  const strokeWidth = props.strokeWidth;

  const previewResource = () => {
    window.open(resourceUrl);
  };

  if (resourceType === "image/*") {
    return (
      <SquareDiv className={clsx(className, "flex items-center justify-center overflow-clip")}>
        <img
          className="min-w-full min-h-full object-cover"
          src={resource.externalLink ? resourceUrl : resourceUrl + "?thumbnail=1"}
          onClick={() => showPreviewImageDialog(resourceUrl)}
          decoding="async"
          loading="lazy"
        />
      </SquareDiv>
    );
  }

  const getResourceIcon = () => {
    switch (resourceType) {
      case "video/*":
        return <Icon.FileVideo2 strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "audio/*":
        return <Icon.FileAudio strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "text/*":
        return <Icon.FileText strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "application/epub+zip":
        return <Icon.Book strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "application/pdf":
        return <Icon.Book strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "application/msword":
        return <Icon.FileEdit strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "application/msexcel":
        return <Icon.SheetIcon strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "application/zip":
        return <Icon.FileArchiveIcon onClick={previewResource} strokeWidth={strokeWidth} className="w-full h-auto" />;
      case "application/x-java-archive":
        return <Icon.BinaryIcon strokeWidth={strokeWidth} className="w-full h-auto" />;
      default:
        return <Icon.File strokeWidth={strokeWidth} className="w-full h-auto" />;
    }
  };

  return (
    <div onClick={previewResource} className={clsx(className, "max-w-[4rem] opacity-50")}>
      {getResourceIcon()}
    </div>
  );
};

export default React.memo(ResourceIcon);
