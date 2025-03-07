import sys
import os

import torch
import soundfile as sf
from transformers import GenerationConfig,WhisperProcessor,AutoModelForSpeechSeq2Seq

# initiate whisper model 
processor = WhisperProcessor.from_pretrained('openai/whisper-base')
model = AutoModelForSpeechSeq2Seq.from_pretrained('mitchelldehaven/whisper-medium-uk')
model.generation_config = GenerationConfig.from_pretrained('openai/whisper-base')

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

def transcribe(file, out):
  audio_input, sample_rate = sf.read(file)

  # todo : read all input in chunks
  chunk = audio_input[0:int(30 * sample_rate)]
  inputs = processor(chunk, sampling_rate=sample_rate, return_tensors='pt', return_attention_mask=True)  
  predicted_ids = model.generate(inputs.input_features, attention_mask=inputs.attention_mask)
  transcription = processor.batch_decode(predicted_ids, skip_special_tokens=True)
  print(transcription)
  
  

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
