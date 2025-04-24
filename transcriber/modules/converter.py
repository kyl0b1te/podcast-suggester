import os
import sys

from pydub import AudioSegment

MODEL_FRAME_RATE = 16000

def get_wav_filename(file_path, out):
  return os.path.join(out, f'{os.path.splitext(os.path.basename(file_path))[0]}.wav')

def get_files_path(src_path, file_name=None):
  files = []
  try:
    for file in os.listdir(src_path):
      is_audio = file.endswith('.mp3') or file.endswith('.m4a')
      is_target = file_name == None or file_name == os.path.basename(file)

      # check file format to skip non audio
      if is_audio and is_target:
        files.append(os.path.join(src_path, file))
  except Exception as e:
    print(f'error: failed to list folder {src_path}:\n{e}')
    sys.exit(1)

  return files

def convert_files(files_path, out):
  print(f'converting {len(files_path)} file(s)')
  for file_path in files_path:
    sound = AudioSegment.from_file(file_path)
    # change frame rate
    sound = sound.set_frame_rate(MODEL_FRAME_RATE)
    # keep only one channel (mono)
    sound = sound.set_channels(1)
    # save wav file
    out_file_path = get_wav_filename(file_path, out)

    sound.export(out_file_path, format='wav')
  print(f'done!\n\nconverted files saved in {out}')

def main(src_path):
  out_path = src_path if os.path.isdir(src_path) else os.path.dirname(src_path)

  files_path = []
  if os.path.isfile(src_path):
    files_path = get_files_path(os.path.dirname(src_path), os.path.basename(src_path))
  elif os.path.isdir(src_path):
    files_path = get_files_path(src_path)
  else:
    print("error: invalid source path parameter")
    sys.exit(1)

  convert_files(files_path, out_path)

