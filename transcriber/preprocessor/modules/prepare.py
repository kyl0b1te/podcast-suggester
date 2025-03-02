import os
import sys

from pydub import AudioSegment

MODEL_FRAME_RATE = 16000 

def get_episodes_path(path):
  files = []
  try:
    for file in os.listdir(path):
      # check file format to skip non audio
      if file.endswith('.mp3') or file.endswith('.m4a'):
        files.append(os.path.join(path, file))
  except Exception as e:
    print(f'error: failed to list folder {path}:\n{e}')
    sys.exit(1)

  return files

def get_wav_filename(path, out):
  return os.path.join(out, f'{os.path.splitext(os.path.basename(path))[0]}.wav')

def preprocess(files, out):
  for file in files:
    filepath = get_wav_filename(file, out)
    sound = AudioSegment.from_file(file)
    # change frame rate
    sound = sound.set_frame_rate(MODEL_FRAME_RATE)
    # keep only one channel (mono)
    sound = sound.set_channels(1)
    # save wav file
    sound.export(get_wav_filename(file, out), format='wav')
    print(f'saved processed file to `{filepath}`')