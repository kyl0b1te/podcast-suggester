import sys
import os

from pydub import AudioSegment 
from pydub.silence import detect_silence

def get_files(path):
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

def get_pauses(filepath):
  sound = AudioSegment.from_file(filepath)
  # todo : tune parameters
  silences = detect_silence(sound, min_silence_len=500, silence_thresh=-16)
  return [(start / 1000.0, end / 1000.0) for start, end in silences]

def transcribe(file, out):
  pauses = get_pauses(file)
  print(pauses)
  

def main(src, ep, out):
  if (src == None and ep == None) or (src != None and ep != None):
      print("error: one of source parameters should be present")
      sys.exit(1)
  
  # prepare list of files for preprocessing
  files = [ep] if ep != None else get_files(src)
  print(f'start process `{len(files)}` file(s)')

  # start preprocessing
  for file in files:
    transcribe(file, out)
