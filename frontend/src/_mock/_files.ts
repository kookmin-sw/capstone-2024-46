import { _mock } from "./_mock";
import { _tags } from "./assets";

// ----------------------------------------------------------------------

const GB = 1000000000 * 24;

const FOLDERS = [
  "Work",
  "Traning1",
  "Traning2",
  "Traning3",
  "Traning4",
  "Traning5",
];

const FILES = ["train1.pdf", "train2.pdf", "train3.pdf", "train4.pdf"];

export const FILE_TYPE_OPTIONS = [
  "folder",
  "txt",
  "zip",
  "audio",
  "image",
  "video",
  "word",
  "excel",
  "powerpoint",
  "pdf",
  "photoshop",
  "illustrator",
];

export const _folders = FOLDERS.map((name, index) => ({
  id: `${_mock.id(index)}_folder`,
  name,
  type: "folder",
  url: "",
  shared: [],
  tags: _tags.slice(0, 5),
  size: GB / ((index + 1) * 10),
  totalFiles: (index + 1) * 100,
  createdAt: _mock.time(index),
  modifiedAt: _mock.time(index),
  isFavorited: _mock.boolean(index + 1),
  isEmbeded: _mock.boolean(index + 1),
}));

export const _files = FILES.map((name, index) => ({
  id: `${_mock.id(index)}_file`,
  name,
  url: "",
  shared: [],
  tags: _tags.slice(0, 5),
  size: GB / ((index + 1) * 500),
  createdAt: _mock.time(index),
  modifiedAt: _mock.time(index),
  type: `${name.split(".").pop()}`,
  isFavorited: _mock.boolean(index + 1),
  isEmbeded: _mock.boolean(index + 1),
}));

export const _allFiles = [..._folders, ..._files];
